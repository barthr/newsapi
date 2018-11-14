package newsapi

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestClient_GetTopHeadlinesWithFailure(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/top-headlines", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		writeJSON(t, w, map[string]interface{}{
			"message": "error",
		})
	})

	response, err := client.GetTopHeadlines(context.Background(), nil)
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

func TestClient_GetTopHeadlinesWithParams(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/top-headlines", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		queryParams := url.Values{}
		queryParams.Add("category", "a")
		queryParams.Add("country", "b")
		queryParams.Add("page", "10")
		testQueryParam(t, r, queryParams)
		writeJSON(t, w, map[string]interface{}{
			"message": "error",
		})
	})

	_, _ = client.GetTopHeadlines(context.Background(), &TopHeadlineParameters{
		Category: "a",
		Country:  "b",
		Page:     10,
	})
}

func TestClient_GetTopHeadlinesWithSuccess(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/top-headlines", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w,
			`
		{

			"status": "ok",
			"totalResults": 1,
			"articles": [
					{
							"source": {
									"id": null,
									"name": "Marketwatch.com"
							},
							"author": "Mark DeCambre",
							"title": "Dow stages 300-point U-turn into negative territory as Apple's stock flirts with bear market",
							"description": null,
							"url": "https://www.marketwatch.com/story/dow-stages-250-point-u-turn-into-negative-territory-as-apples-stock-flirts-with-bear-market-2018-11-14",
							"urlToImage": "https://mw3.wsj.net/mw5/content/logos/mw_logo_social.png",
							"publishedAt": "2018-11-14T16:25:00Z",
							"content": "The Dow Jones Industrial Average gave up a more than 200-point opening gain early Wednesday and turned negative, as shares of Apple Inc. AAPL, -2.08% slumped. The Dow DJIA, -0.49% had opened the session with an advance of as many as 214 points but was most re… [+819 chars]"
					}
			]
		}
		`)
	})

	response, err := client.GetTopHeadlines(context.Background(), nil)
	if err != nil {
		t.Errorf("expected error to be equals to nil got %v", err)
	}
	if response == nil {
		t.Error("expected response to be equal to something")
	}
	if len(response.Articles) != 1 {
		t.Errorf("expected 1 source bot got %d", len(response.Articles))
	}
	if response.TotalResults != 1 {
		t.Errorf("expected totalresults to be equal to 1 but got %v", response.TotalResults)
	}
}
