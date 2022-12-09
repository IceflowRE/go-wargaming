package wargaming

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// response Wargaming response body fits always in this struct.
type response struct {
	Status string         `json:"status"`
	Error  *ResponseError `json:"error,omitempty"`
	Data   any            `json:"data,omitempty"`
	Meta   any            `json:"meta,omitempty"`
}

// Client for all api requests.
// As a quick start:
//
//	// create client
//	client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", nil)
//	// get account list
//	res, err := client.Wot.AccountList(context.Background(), wargaming.realmEu, "Yzne", nil)
//	// print out
//	if err != nil {
//		// handle a Wargaming Response Error, returned by the API itself
//		var respErr *wargaming.ResponseError
//		if errors.As(err, &respErr) {
//			fmt.Println(respErr.Error())
//		} else {
//			// handle client or other error
//			fmt.Println(err.Error())
//		}
//	} else {
//		for _, value := range res {
//			fmt.Println(value)
//		}
//	}
//	// get account list with options
//	res, err = client.Wot.AccountList(context.Background(), wargaming.realmEu, "Lichtgeschwindigkeit", &wot.AccountListOptions{
//		Fields: []string{"nickname"},
//	})
//
// or with a custom http.Client
//
//		client := wargaming.NewClient("a7f838650dcb008552966db063eeeb35", &wargaming.ClientOptions{
//	   HTTPClient: &http.Client{
//	       Timeout: 10 * time.Second,
//	   },
//		})
type Client struct {
	wgServices
	httpClient *http.Client
	// Wargaming.net application ID.
	applicationId string

	// Reuse a single struct instead of allocating one for each service on the heap.
	common service
}

type ClientOptions struct {
	HTTPClient *http.Client
}

// NewClient creates a new API client.
// Pass nil as options if you want to use none.
func NewClient(applicationId string, options *ClientOptions) *Client {
	client := &Client{
		applicationId: applicationId,
	}
	if options != nil {
		if options.HTTPClient != nil {
			client.httpClient = options.HTTPClient
		} else {
			client.httpClient = &http.Client{}
		}
	}
	client.common.client = client
	client.wgServices = newWgServices(&client.common)
	return client
}

func (client *Client) buildRequest(ctx context.Context, method string, section section, realm Realm, path string, data map[string]string) (req *http.Request) {
	query := url.Values{}
	query.Add("application_id", client.applicationId)
	for key, value := range data {
		query.Add(key, value)
	}
	if method == http.MethodPost {
		req, _ = http.NewRequestWithContext(ctx, method, section.ApiUrl(realm)+path, bytes.NewBuffer([]byte(query.Encode())))
	} else {
		req, _ = http.NewRequestWithContext(ctx, method, section.ApiUrl(realm)+path, nil)
		req.URL.RawQuery = query.Encode()
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}
	return req
}

func (client *Client) request(req *http.Request, returnData any, metaData any) error {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return BadStatusCode(resp.StatusCode)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	wgResp := &response{
		Status: "",
		Error:  nil,
		Data:   returnData,
		Meta:   metaData,
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

// postRequest set returnData to nil, if no response data is expected.
// set metaData to nil if no meta data is expected.
func (client *Client) postRequest(ctx context.Context, section section, realm Realm, path string, data map[string]string, returnData any, metaData any) error {
	req := client.buildRequest(ctx, "POST", section, realm, path, data)
	return client.request(req, returnData, metaData)
}

// getRequest set returnData to nil, if no response data is expected.
// set metaData to nil if no meta data is expected.
func (client *Client) getRequest(ctx context.Context, section section, realm Realm, path string, data map[string]string, returnData any, metaData any) error {
	req := client.buildRequest(ctx, "GET", section, realm, path, data)
	return client.request(req, returnData, metaData)
}
