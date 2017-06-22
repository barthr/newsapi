package newsapi

import "net/http"

type SourceParameters struct {
	Category string `url:"category,omitempty"`
	Language string `url:"language,omitempty"`
	Country  string `url:"country,omitempty"`
}

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

type SourceResponse struct {
	Status  string   `json:"status"`
	Sources []Source `json:"sources"`
}

func (c *Client) GetSources(params *SourceParameters) (*SourceResponse, *http.Response, error) {
	u := "sources"

	if params != nil {
		var err error
		u, err = setOptions(u, params)

		if err != nil {
			return nil, nil, err
		}
	}

	req, err := c.NewGetRequest(u)
	if err != nil {
		return nil, nil, err
	}

	var response *SourceResponse

	resp, err := c.Do(req, &response)

	if err != nil {
		return nil, nil, err
	}

	return response, resp, nil
}
