package newsapi

import (
	"encoding/json"
	"io"
	"io/ioutil"
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
	BaseURL *url.URL
	// Api Key used with requests to NewsAPI.
	apiKey string
	// User agent used when communicating with the NewsAPI API.
	UserAgent string
}

func NewClient(apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURL, _ := url.Parse(defaultBaseURL)

	return &Client{
		client:  httpClient,
		apiKey:  apiKey,
		BaseURL: baseURL,
	}
}

func (c *Client) NewGetRequest(urlStr string) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set(apiKeyHeader, c.apiKey)

	if c.UserAgent != "" {
		req.Header.Set("User-Agent", c.UserAgent)
	}
	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)

	if err != nil {
		return resp, err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		// Found in https://github.com/google/go-github
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	err = json.NewDecoder(resp.Body).Decode(v)
	if err == io.EOF {
		err = nil
	}

	return resp, err
}

func setOptions(reqUrl string, opts interface{}) (string, error) {
	u, err := url.Parse(reqUrl)
	if err != nil {
		return reqUrl, err
	}

	qs, err := query.Values(opts)
	if err != nil {
		return reqUrl, err
	}

	u.RawQuery = qs.Encode()
	return u.String(), nil
}
