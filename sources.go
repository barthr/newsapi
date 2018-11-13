package newsapi

import (
	"context"
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
}

// SourceResponse is the response from the source request
type SourceResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

// GetSources returns the sources from newsapi see http://newsapi.org/docs for more information on the parameters
func (c *Client) GetSources(ctx context.Context, params *SourceParameters) (*SourceResponse, error) {
	u := sourcesEndpoint

	if params != nil {
		var err error
		u, err = setOptions(u, params)

		if err != nil {
			return nil, err
		}
	}

	req, err := c.newGetRequest(u)
	if err != nil {
		return nil, err
	}

	var response = struct {
		*Error
		*SourceResponse
	}{}

	_, err = c.do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.SourceResponse, nil
}
