package newsapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_GetSourcesWithFailure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(t, w, map[string]interface{}{
			"message": "error",
		})
	})

	response, err := client.GetSources(context.Background(), nil)
	if err == nil {
		t.Error("expected error to be not equals to nil")
	}
	if !APIError(err) {
		t.Errorf("expected error to be an api error but got %v", err)
	}
	if err.(*Error).Message != "error" {
		t.Errorf("expected message to be equal to error but got %s", err.(*Error).Message)
	}
	if response != nil {
		t.Error("expected response to be equal to nil")
	}
}

func TestClient_GetSourcesWithParams(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		queryParams := url.Values{}
		queryParams.Add("category", "a")
		queryParams.Add("country", "b")
		queryParams.Add("language", "c")
		testQueryParam(t, r, queryParams)
		writeJSON(t, w, map[string]interface{}{
			"message": "error",
		})
	})

	_, _ = client.GetSources(context.Background(), &SourceParameters{
		Category: "a",
		Country:  "b",
		Language: "c",
	})
}

func TestClient_GetSourcesWithResponse(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/sources", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `
		{

			"status": "ok",
			"sources": [
					{
							"id": "abc-news",
							"name": "ABC News",
							"description": "Your trusted source for breaking news, analysis, exclusive interviews, headlines, and videos at ABCNews.com.",
							"url": "https://abcnews.go.com",
							"category": "general",
							"language": "en",
							"country": "us"
					}
			]
		}`)
	})

	response, err := client.GetSources(context.Background(), nil)
	if err != nil {
		t.Errorf("expected error to be equals to nil got %v", err)
	}
	if response == nil {
		t.Error("expected response to be equal to something")
	}
	if len(response.Sources) != 1 {
		t.Errorf("Expected 1 source bot got %d", len(response.Sources))
	}
}

func TestClient_GetSourcesWithFailingHTTPClient(t *testing.T) {
	client, _, teardown := setup()
	defer teardown()

	response, err := client.GetSources(context.Background(), nil)
	if err == nil {
		t.Error("expected error to be equals to not nil got nil")
	}
	if response != nil {
		t.Error("expected response to be equal to something")
	}
}
