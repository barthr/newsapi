<p align="center">
    <img src ="logo.png"></img>
</p>

# NewsAPI Go Client

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/barthr/newsapi) [![Build Status](https://travis-ci.org/barthr/newsapi.svg?branch=master)](https://travis-ci.org/barthr/newsapi) [![codecov](https://codecov.io/gh/barthr/newsapi/branch/master/graph/badge.svg)](https://codecov.io/gh/barthr/newsapi)
[![Golangci](https://golangci.com/badges/github.com/barthr/newsapi.svg)](https://golangci.com/r/github.com/barthr/newsapi)

   
Go client for communicating with the newsapi api.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

```
go get github.com/barthr/newsapi
```

Next up register for free at (https://newsapi.org/register) get yourself a free api key and keep it somewhere save.


## Examples

### Retrieving all the sources

```go
package main

import (
	"fmt"
	"net/http"
	"context"
	"github.com/barthr/newsapi"
)

func main() {
	c := newsapi.NewClient("<API KEY>", newsapi.WithHTTPClient(http.DefaultClient))

	sources, err := c.GetSources(context.Background(), nil)

	if err != nil {
		panic(err)
	}

	for _, s := range sources.Sources {
		fmt.Println(s.Description)
	}
}
```

### Retrieving all the sources for a specific country (Great britian in this case)

```go
package main

import (
	"fmt"
	"net/http"
	"context"

	"github.com/barthr/newsapi"
)

func main() {
	c := newsapi.NewClient("<API KEY>", newsapi.WithHTTPClient(http.DefaultClient))

	sources, err := c.GetSources(context.Background(), &newsapi.SourceParameters{
		Country: "gb",
	})

	if err != nil {
		panic(err)
	}

	for _, s := range sources.Sources {
		fmt.Println(s.Name)
	}
}
```

### Retrieving the top headlines

```go
package main

import (
	"fmt"
	"net/http"
	"context"

	"github.com/barthr/newsapi"
)

func main() {
	c := newsapi.NewClient("<API KEY>", newsapi.WithHTTPClient(http.DefaultClient))

	articles, err := c.GetTopHeadlines(context.Background(), &newsapi.TopHeadlineParameters{
		Sources: []string{ "cnn", "time" },
	})

	if err != nil {
		panic(err)
	}

	for _, s := range articles.Articles {
		fmt.Printf("%+v\n\n", s)
	}
}
```

### Retrieving all the articles

```go
package main

import (
	"fmt"
	"net/http"
	"context"

	"github.com/barthr/newsapi"
)

func main() {
	c := newsapi.NewClient("<API KEY>", newsapi.WithHTTPClient(http.DefaultClient))

	articles, err := c.GetEverything(context.Background(), &newsapi.EverythingParameters{
		Sources: []string{ "cnn", "time" },
	})

	if err != nil {
		panic(err)
	}

	for _, s := range articles.Articles {
		fmt.Printf("%+v\n\n", s)
	}
}
```


## License

This project is licensed under the MIT License

## Acknowledgments

* Inspiration from github golang client
