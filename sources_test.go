package newsapi

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestGetSources(t *testing.T) {
	var tests = map[string]struct {
		client   *Client
		ctx      context.Context
		params   *SourceParameters
		response *SourceResponse
		wantErr  bool
	}{
		"test with no parameters": {
			client: func() *Client {
				transport := httpmock.NewMockTransport()
				transport.RegisterResponder(http.MethodGet, "https://newsapi.org/v2/sources", nil)

				httpClient := &http.Client{
					Transport: transport,
				}
				return NewClient("", WithHTTPClient(httpClient))
			}(),
			ctx:     context.Background(),
			params:  nil,
			wantErr: true,
		},
		"test failure from api": {
			client: func() *Client {
				transport := httpmock.NewMockTransport()
				transport.RegisterResponder(http.MethodGet, "https://newsapi.org/v2/sources", func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("failure")
				})

				httpClient := &http.Client{
					Transport: transport,
				}
				return NewClient("", WithHTTPClient(httpClient))
			}(),
			ctx:     context.Background(),
			params:  nil,
			wantErr: true,
		},
		"test with setting source params": {
			client: func() *Client {
				transport := httpmock.NewMockTransport()
				transport.RegisterResponder(http.MethodGet, "https://newsapi.org/v2/sources", func(req *http.Request) (*http.Response, error) {
					return httpmock.NewStringResponse(
						http.StatusOK,
						`
						{
							"status": "ok",
							"sources": [
								{

									"id": "abc-news-au",
									"name": "ABC News (AU)",
									"description": "Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.",
									"url": "http://www.abc.net.au/news",
									"category": "general",
									"language": "en",
									"country": "au"
								}
							]
						}
						`,
					), nil
				})

				httpClient := &http.Client{
					Transport: transport,
				}
				return NewClient("", WithHTTPClient(httpClient))
			}(),
			ctx: context.Background(),
			params: &SourceParameters{
				Category: "general",
				Country:  "au",
				Language: "en",
			},
			response: &SourceResponse{
				Status: "ok",
				Sources: []Source{
					Source{
						Category:    "general",
						ID:          "abc-news-au",
						Name:        "ABC News (AU)",
						URL:         "http://www.abc.net.au/news",
						Description: "Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.",
						Language:    "en",
						Country:     "au",
					},
				},
			},
			wantErr: false,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			response, err := test.client.GetSources(test.ctx, test.params)
			if !reflect.DeepEqual(test.response, response) {
				t.Errorf("Got %v not equal to expected %v", response, test.response)
			}
			if (err != nil) != test.wantErr {
				t.Errorf("Expected (err != nil) %v to be %v", err != nil, test.wantErr)
			}
		})
	}
}
