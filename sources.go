package newsapi

import (
	"context"
	"net/http"
)

// SourceParameters are the parameters which can be used in the source request to newsapi
type SourceParameters struct {
	Category string `url:"category,omitempty"`
	Language string `url:"language,omitempty"`
	Country  string `url:"country,omitempty"`
}

// Source is a source from newsapi it contains all the information provided by the api
type Source struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	Language    string `json:"language"`
	Country     string `json:"country"`
	UrlsToLogos struct {
		Small  string `json:"small"`
		Medium string `json:"medium"`
		Large  string `json:"large"`
	} `json:"urlsToLogos"`
	SortBysAvailable []string `json:"sortBysAvailable"`
}

// SourceResponse is the response from the source request
type SourceResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

// GetSources returns the sources from newsapi see https://newsapi.org/#apiSources for more information on the parameters
func (c *Client) GetSources(ctx context.Context, params *SourceParameters) (*SourceResponse, *http.Response, error) {
	u := "sources"

	if params != nil {
		var err error
		u, err = setOptions(u, params)

		if err != nil {
			return nil, nil, err
		}
	}

	req, err := c.newGetRequest(u)
	if err != nil {
		return nil, nil, err
	}

	var response *SourceResponse

	resp, err := c.do(ctx, req, &response)

	if err != nil {
		return nil, nil, err
	}

	return response, resp, nil
}
