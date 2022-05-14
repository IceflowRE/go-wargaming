package wargaming

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error,omitempty"`
	Data   any            `json:"data,omitempty"`
}

type Client struct {
	httpClient    *http.Client
	applicationId string
	// Do a local parameter check before actually sending the request.
	// Disable if the library might be outdated and has not all current available parameters.
	localCheck bool
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

func (client *Client) WithLocalCheck() *Client {
	client.localCheck = true
	return client
}

func (client *Client) WithHttpClient(httpClient *http.Client) *Client {
	client.httpClient = httpClient
	return client
}

// returnData must be a pointer
func (client *Client) sendGetRequest(realm Realm, path string, data map[string]string, returnData any) error {
	req, _ := http.NewRequest("GET", realm.ApiUrl()+"/"+path, nil)
	query := req.URL.Query()
	query.Add("application_id", client.applicationId)
	for key, value := range data {
		query.Add(key, value)
	}
	req.URL.RawQuery = query.Encode()

	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return BadStatusCode(resp.StatusCode)
	}

	// read into json-RawMessage map
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	wgResp := &Response{
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
