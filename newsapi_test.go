package newsapi

import (
	"net/http"
	"net/http/httptest"
	"net/url"
)

// setup sets up a test HTTP server along with a github.Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, serverURL string, teardown func()) {
	// mux is the HTTP request multiplexer used with the test server.
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(defaultBaseURL+"/", http.StripPrefix(defaultBaseURL, mux))

	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	// client is the GitHub client being tested and is
	// configured to use test server.
	url, _ := url.Parse(server.URL + defaultBaseURL + "/")
	client = NewClient("", WithBaseURL(url))

	return client, mux, server.URL, server.Close
}
