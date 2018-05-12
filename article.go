package newsapi

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"
)

// ArticleParameters are the parameters used for the newsapi article endpoint
// Source must always contain a value
// See http://beta.newsapi.org/docs for more information on the required parameters
type ArticleParameters struct {
	Sources []string `url:"sources,omitempty,comma"`
	Domains []string `url:"domains,omitempty,comma"`

	Keywords string `url:"q,omitempty"`
	Category string `url:"category,omitempty"`
	Language string `url:"language,omitempty"`
	SortBy   string `url:"sortBy,omitempty"`
	Page     int    `url:"page,omitempty"`
}

// Article is a single article from the newsapi article response
// See http://beta.newsapi.org/docs for more details on the property's
type Article struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
}

// ArticleResponse is the response from the newsapi article endpoint
// Code and Message property will be filled when an error happened
// See http://beta.newsapi.org/docs for more details on the property's
type ArticleResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Code         string    `json:"code,omitempty"`
	Message      string    `json:"message,omitempty"`
	Articles     []Article `json:"articles"`
}

// GetTopHeadlines returns the articles from newsapi
// See http://beta.newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) GetTopHeadlines(ctx context.Context, params *ArticleParameters) (*ArticleResponse, *http.Response, error) {
	return c.getArticles(ctx, "top-headlines", params)
}

// GetEverything returns the articles from newsapi
// See http://beta.newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) GetEverything(ctx context.Context, params *ArticleParameters) (*ArticleResponse, *http.Response, error) {
	return c.getArticles(ctx, "everything", params)
}

// GetArticles returns the articles from newsapi
// See http://beta.newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) getArticles(ctx context.Context, u string, params *ArticleParameters) (*ArticleResponse, *http.Response, error) {
	if params == nil {
		return nil, nil, errors.New("empty parameters not possible when asking for articles")
	}

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

	var response *ArticleResponse

	resp, err := c.do(ctx, req, &response)

	if err != nil {
		return nil, nil, err
	}

	if response.Code != "" {
		return nil, nil, fmt.Errorf("[%s] %s", response.Code, response.Message)
	}

	return response, resp, nil
}
