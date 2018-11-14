// Package newsapi provides helper functions to query https://newsapi.org
// The api maps the responses to appropriate go structs and include all of the possible options
package newsapi

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	apiKeyHeader    = "X-Api-Key"
	userAgentHeader = "User-Agent"
	defaultBaseURL  = "https://newsapi.org/v2/"

	sourcesEndpoint      = "sources"
	topHeadlinesEndpoint = "top-headlines"
	everythingEndpoint   = "everything"
)

// A Client manages communication with the NewsAPI API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.
	// Base URL for API requests
	baseURL *url.URL
	// Api Key used with requests to NewsAPI.
	apiKey string
	// User agent used when communicating with the NewsAPI API.
	userAgent string
}

// OptionFunc is function which modifies the client
type OptionFunc func(*Client)

// WithHTTPClient sets the http client to use when making requests.
func WithHTTPClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.client = client
	}
}

// WithBaseURL sets the baseurl for the newsapi
func WithBaseURL(url *url.URL) OptionFunc {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithUserAgent sets the user agent of the client to userAgent
func WithUserAgent(userAgent string) OptionFunc {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// NewClient returns a new newsapi client to query the newsapi API.
func NewClient(apiKey string, options ...OptionFunc) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		baseURL:   baseURL,
		apiKey:    apiKey,
		userAgent: "github.com/barthr/newsapi",
	}

	for _, opt := range options {
		opt(c)
	}
	return c
}

// newGetRequest returns a new Get request for the given url URLStr
// It returns a pointer to a http request which can be executed by a http.client
func (c *Client) newGetRequest(URLStr string) (*http.Request, error) {
	rel, err := url.Parse(URLStr)
	if err != nil {
		return nil, err
	}

	u := c.baseURL.ResolveReference(rel)

	// Create a new get request with the url provided
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Set the api key on the request
	req.Header.Set(apiKeyHeader, c.apiKey)

	// If we specify a user agent we override the current one
	if c.userAgent != "" {
		req.Header.Set(userAgentHeader, c.userAgent)
	}
	return req, nil
}

// do executes the http.Request and marshal's the response into v
// v must be a pointer to a value instead of a regular value
// It returns the actual response from the request and also an error
func (c *Client) do(ctx context.Context, req *http.Request, v interface{}) (*http.Response, error) {
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)

	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err == io.EOF {
		err = nil
	}

	return resp, err
}

// setOptions set the options for the url
// It will set the query parameters and encodes them with the url
func setOptions(reqURL string, options interface{}) (string, error) {
	u, err := url.Parse(reqURL)
	if err != nil {
		return reqURL, err
	}

	qs, err := query.Values(options)
	if err != nil {
		return reqURL, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
