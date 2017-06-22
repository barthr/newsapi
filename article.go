package newsapi

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type ArticleParameters struct {
	Source string `url:"source,omitempty"`
	SortBy string `url:"sortBy,omitempty"`
}

type Article struct {
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

type ArticleResponse struct {
	Status   string    `json:"status"`
	Code     string    `json:"code,omitempty"`
	Message  string    `json:"message,omitempty"`
	Source   string    `json:"source"`
	SortBy   string    `json:"sortBy"`
	Articles []Article `json:"articles"`
}

func (c *Client) GetArticles(params *ArticleParameters) (*ArticleResponse, *http.Response, error) {
	u := "articles"

	if params == nil || params.Source == "" {
		return nil, nil, errors.New("empty source not possible when asking for articles")
	}

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

	var response *ArticleResponse

	resp, err := c.Do(req, &response)

	if err != nil {
		return nil, nil, err
	}

	if response.Code != "" {
		return nil, nil, fmt.Errorf("[%s] %s", response.Code, response.Message)
	}

	return response, resp, nil
}
