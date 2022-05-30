package wargaming

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error,omitempty"`
	Data   any            `json:"data,omitempty"`
}

// Client for all api requests.
// As a quick start:
//	// create client
//	client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35")
//	// get account list
//	res, err := client.WotAccountList(wargaming.RealmEu, "Yzne", []string{}, "", 0, "")
//	// print out
//	if err != nil {
//		fmt.Println(err.Error())
//	} else {
//		for _, value := range res {
//			fmt.Println(value)
//		}
//	}
// or with a custom http.Client
//	client := wargaming.NewClient("d5dcedcbea865ea892202ba7361e626c").
//		WithHttpClient(&http.Client{
//		Timeout: 10 * time.Second,
//	})
type Client struct {
	httpClient    *http.Client
	applicationId string
}

func NewClient(applicationId string) *Client {
	return &Client{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConns:    20,
				IdleConnTimeout: 5 * time.Second,
			},
		},
		applicationId: applicationId,
	}
}

func (client *Client) WithHttpClient(httpClient *http.Client) *Client {
	client.httpClient = httpClient
	return client
}

func (client *Client) buildRequest(realm Realm, path string, data map[string]string, method string) (req *http.Request) {
	query := url.Values{}
	query.Add("application_id", client.applicationId)
	for key, value := range data {
		query.Add(key, value)
	}
	if method == "POST" {
		req, _ = http.NewRequest(method, realm.ApiUrl()+path, bytes.NewBuffer([]byte(query.Encode())))
	} else {
		req, _ = http.NewRequest(method, realm.ApiUrl()+path, nil)
		req.URL.RawQuery = query.Encode()
	}
	return req
}

func (client *Client) doRequest(req *http.Request, returnData any) error {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return BadStatusCode(resp.StatusCode)
	}

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	wgResp := &response{
		Status: "",
		Error:  nil,
		Data:   returnData,
	}
	err = json.Unmarshal(respBytes, wgResp)
	if err != nil {
		return err
	}

	if wgResp.Error != nil {
		return wgResp.Error
	}
	if wgResp.Status != "ok" {
		return InvalidResponse
	}
	return nil
}

func (client *Client) doPostRequest(realm Realm, path string, data map[string]string) error {
	req := client.buildRequest(realm, path, data, "POST")
	return client.doRequest(req, struct{}{})
}

func (client *Client) doGetRequest(realm Realm, path string, data map[string]string) error {
	req := client.buildRequest(realm, path, data, "GET")
	return client.doRequest(req, struct{}{})
}

// returnData must be a pointer
func (client *Client) doPostDataRequest(realm Realm, path string, data map[string]string, returnData any) error {
	req := client.buildRequest(realm, path, data, "POST")
	return client.doRequest(req, returnData)
}

// returnData must be a pointer
func (client *Client) doGetDataRequest(realm Realm, path string, data map[string]string, returnData any) error {
	req := client.buildRequest(realm, path, data, "GET")
	return client.doRequest(req, returnData)
}
