package newsapi

import (
	"context"
	"time"
)

// TopHeadlineParameters are the parameters which can be used to tweak to request for the top headlines.
type TopHeadlineParameters struct {
	Country  string   `url:"country,omitempty"`
	Category string   `url:"category,omitempty"`
	Sources  []string `url:"sources,omitempty,comma"`
	Keywords string   `url:"q,omitempty"`
	Page     int      `url:"page,omitempty"`
	PageSize int      `url:"pageSize,omitempty"`
}

// EverythingParameters are the parameters used for the newsapi everything endpoint.
type EverythingParameters struct {
	Keywords       string   `url:"q,omitempty"`
	Sources        []string `url:"sources,omitempty,comma"`
	Domains        []string `url:"domains,omitempty,comma"`
	ExcludeDomains []string `url:"excludeDomains,omitempty"`

	From time.Time `url:"from,omitempty"`
	To   time.Time `url:"to,omitempty"`

	Language string `url:"language,omitempty"`
	SortBy   string `url:"sortBy,omitempty"`

	Page     int `url:"page,omitempty"`
	PageSize int `url:"pageSize,omitempty"`
}

// Article is a single article from the newsapi article response
// See http://newsapi.org/docs for more details on the property's
type Article struct {
	Source      Source    `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

// ArticleResponse is the response from the newsapi article endpoint.
// Code and Message property will be filled when an error happened.
// See http://newsapi.org/docs for more details on the property's.
type ArticleResponse struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

// GetTopHeadlines returns the articles from newsapi
// See http://newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) GetTopHeadlines(ctx context.Context, params *TopHeadlineParameters) (*ArticleResponse, error) {
	return c.getArticles(ctx, topHeadlinesEndpoint, params)
}

// GetEverything returns the articles from newsapi
// See http://newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) GetEverything(ctx context.Context, params *EverythingParameters) (*ArticleResponse, error) {
	return c.getArticles(ctx, everythingEndpoint, params)
}

// GetArticles returns the articles from newsapi
// See http://newsapi.org/docs for more information
// It will return the error from newsapi if there is an error
func (c *Client) getArticles(ctx context.Context, u string, params interface{}) (*ArticleResponse, error) {
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
		*ArticleResponse
	}{}

	_, err = c.do(ctx, req, &response)

	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return response.ArticleResponse, nil
}
