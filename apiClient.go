/*
Package apiClient provides a reusable library for sending external API requests with features like debugging, retries, and request/response logging.

Author: Mohamed Riyad
Email: m@ryad.dev
Website: https://ryad.dev
*/

package apiClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Client represents the API client.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
	Debug      bool
	MaxRetries int
	RetryDelay time.Duration
}

// NewClient initializes and returns a new Client.
//
// baseURL: the base URL for the API.
// debug: enables or disables debug logging.
// maxRetries: the maximum number of retries for failed requests.
// retryDelay: the delay between retries.
func NewClient(baseURL string, debug bool, maxRetries int, retryDelay time.Duration) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Debug:      debug,
		MaxRetries: maxRetries,
		RetryDelay: retryDelay,
	}
}

// APIRequest represents the structure of an API request.
type APIRequest struct {
	Method   string
	Endpoint string
	Headers  map[string]string
	Body     interface{}
}

// APIResponse represents the structure of an API response.
type APIResponse struct {
	StatusCode int
	Headers    map[string][]string
	Body       []byte
}

// SendRequest sends an HTTP request and returns the response.
//
// req: the APIRequest object containing request details.
// returns: an APIResponse object or an error if the request fails.
func (c *Client) SendRequest(req *APIRequest) (*APIResponse, error) {
	var response *APIResponse
	var err error

	for attempt := 0; attempt <= c.MaxRetries; attempt++ {
		response, err = c.send(req)
		if err == nil {
			return response, nil
		}
		time.Sleep(c.RetryDelay)
	}

	return response, err
}

// send sends the HTTP request and handles the response.
//
// req: the APIRequest object containing request details.
// returns: an APIResponse object or an error if the request fails.
func (c *Client) send(req *APIRequest) (*APIResponse, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, req.Endpoint)

	var reqBody []byte
	var err error
	if req.Body != nil {
		reqBody, err = json.Marshal(req.Body)
		if err != nil {
			return nil, err
		}
	}

	request, err := http.NewRequest(req.Method, url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	for key, value := range req.Headers {
		request.Header.Set(key, value)
	}

	if c.Debug {
		c.logRequest(request, reqBody)
	}

	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		c.logResponse(response, respBody)
	}

	apiResponse := &APIResponse{
		StatusCode: response.StatusCode,
		Headers:    response.Header,
		Body:       respBody,
	}

	return apiResponse, nil
}

// logRequest logs the details of an HTTP request.
//
// request: the HTTP request object.
// body: the request body as a byte slice.
func (c *Client) logRequest(request *http.Request, body []byte) {
	fmt.Printf("Request Method: %s\n", request.Method)
	fmt.Printf("Request URL: %s\n", request.URL.String())
	fmt.Printf("Request Headers: %v\n", request.Header)
	fmt.Printf("Request Body: %s\n", string(body))
}

// logResponse logs the details of an HTTP response.
//
// response: the HTTP response object.
// body: the response body as a byte slice.
func (c *Client) logResponse(response *http.Response, body []byte) {
	fmt.Printf("Response Status: %s\n", response.Status)
	fmt.Printf("Response Headers: %v\n", response.Header)
	fmt.Printf("Response Body: %s\n", string(body))
}
