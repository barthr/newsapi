package newsapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

// setup sets up a test HTTP server along with a Client that is
// configured to talk to that test server. Tests should register handlers on
// mux which provide mock responses for the API method being tested.
func setup() (client *Client, mux *http.ServeMux, teardown func()) {
	apiHandler := http.NewServeMux()
	// server is a test HTTP server used to provide mock API responses.
	server := httptest.NewServer(apiHandler)

	url, _ := url.Parse(server.URL + "/")
	client = NewClient("", WithBaseURL(url))

	return client, apiHandler, server.Close
}

func writeJSON(t *testing.T, w http.ResponseWriter, v interface{}) {
	if err := json.NewEncoder(w).Encode(v); err != nil {
		t.Errorf("failed marshalling v to response writer because %v", v)
	}
}

func testQueryParam(t *testing.T, r *http.Request, want url.Values) {
	if got := r.URL.Query(); !reflect.DeepEqual(got, want) {
		t.Errorf("expected url values aren't the same. got %v want %v", got, want)
	}
}

func TestClient_WithUserAgent(t *testing.T) {
	userAgent := "testing"
	client := NewClient("", WithUserAgent(userAgent))
	if client.userAgent != userAgent {
		t.Errorf("expected user agent to be set to %s but got %s", userAgent, client.userAgent)
	}
}

func TestClient_WithHTTPClient(t *testing.T) {
	client := NewClient("", WithHTTPClient(nil))
	if client.client != nil {
		t.Errorf("expected http client set to nil but got %v", client.client)
	}
}

func TestNewGetRequestHeaders(t *testing.T) {
	userAgent := "hello"
	apiKey := "goodbye"
	c := NewClient(apiKey, WithUserAgent(userAgent))
	req, err := c.newGetRequest("$^_dfa0s9dfas/failure###/$$/##https://")
	if err != nil {
		t.Errorf("expected error to be equal to nil but got %v", err)
	}
	if req == nil {
		t.Error("expected request to be created")
	}
	if req.Header.Get("User-Agent") != userAgent {
		t.Errorf("expected user agent to be equal to %s but got %s", userAgent, req.Header.Get("User-Agent"))
	}
	if req.Header.Get(apiKeyHeader) != apiKey {
		t.Errorf("expected api key header to be equal to %s but got %s", apiKey, req.Header.Get(apiKeyHeader))
	}
}
