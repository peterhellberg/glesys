package glesys

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var (
	ErrUnexpectedHTTPStatus = fmt.Errorf("unexpected HTTP status")
)

// A Client communicates with the GleSYS API.
type Client struct {
	// Username is the username used for requests to the GleSYS API
	Username string

	// APIKey is the key used for requests to the GleSYS API
	APIKey string

	// BaseURL is the base url for GleSYS API.
	BaseURL *url.URL

	// BasePath is the base path for the gifs endpoints
	BasePath string

	// User agent used for HTTP requests to GleSYS API.
	UserAgent string

	// HTTP client used to communicate with the GleSYS API.
	httpClient *http.Client
}

// Status contains the status fields for responses from the GleSYS API
type Status struct {
	Code      int       `json:"code"`
	Timestamp time.Time `json:"timestamp"`
	Text      string    `json:"text"`
}

// Debug contains the input variables sent to the GleSYS API
type Debug struct {
	Input map[string]interface{} `json:"input"`
}

// NewClient returns a new GleSYS API client.
// If httpClient is nil, http.DefaultClient is used.
func NewClient(username, apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		cloned := *http.DefaultClient
		httpClient = &cloned
	}

	c := &Client{
		Username: username,
		APIKey:   apiKey,
		BaseURL: &url.URL{
			Scheme: "https",
			Host:   "api.glesys.com",
		},
		UserAgent:  "glesys.go",
		httpClient: httpClient,
	}

	return c
}

// NewRequest creates an API request.
func (c *Client) NewRequest(method, path string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse("/" + path)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.Username, c.APIKey)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", c.UserAgent)

	if method == "POST" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// GetRequest creates a new GET request
func (c *Client) GetRequest(path string) (*http.Request, error) {
	return c.NewRequest("GET", path, nil)
}

// PostRequest creates a new POST request
func (c *Client) PostRequest(path string, data url.Values) (*http.Request, error) {
	return c.NewRequest("POST", path, strings.NewReader(data.Encode()))
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	// Make sure to close the connection after replying to this request
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return resp, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s %s: %s", req.Method, req.URL.RequestURI(), err)
	}

	return resp, nil
}
