package newsapi

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/google/go-querystring/query"
)

const (
	apiKeyHeader   = "X-Api-Key"
	defaultBaseURL = "https://newsapi.org/v1/"
)

// A Client manages communication with the NewsAPI API.
type Client struct {
	client *http.Client // HTTP client used to communicate with the API.
	// Base URL for API requests
	baseURL *url.URL
	// Api Key used with requests to NewsAPI.
	apiKey string
	// User agent used when communicating with the NewsAPI API.
	UserAgent string
}

type OptionFunc func(*Client)

func WithHttpClient(client *http.Client) OptionFunc {
	return func(c *Client) {
		c.client = client
	}
}

// NewClient returns a new newsapi client to query the newsapi API.
func NewClient(apiKey string, options ...OptionFunc) *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{
		client:  http.DefaultClient,
		baseURL: baseURL,
		apiKey:  apiKey,
	}

	for _, opt := range options {
		opt(c)
	}
	return c
}

// NewGetRequest returns a new Get request for the given url URLStr
// It returns a pointer to a http request which can be executed by a http.client
func (c *Client) NewGetRequest(URLStr string) (*http.Request, error) {
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
	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

// Do executes the http.Request and marshal's the response into v
// v must be a pointer to a value instead of a regular value
// It returns the actual response from the request and also an error
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
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
