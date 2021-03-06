package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/barthr/newsapi"
)

func getNewsSources() {
	news_api_key := os.Getenv("NEWS_API_KEY")

	c := newsapi.NewClient(news_api_key, newsapi.WithHTTPClient(http.DefaultClient))

	sources, err := c.GetSources(context.Background(), &newsapi.SourceParameters{
		Language: "en",
	})

	if err != nil {
		panic(err)
	}

	for _, s := range sources.Sources {
		fmt.Println(s.Name)

	}
}

func getNewsArticles() {
	news_api_key := os.Getenv("NEWS_API_KEY")

	c := newsapi.NewClient(news_api_key, newsapi.WithHTTPClient(http.DefaultClient))

	articles, err := c.GetEverything(context.Background(), &newsapi.EverythingParameters{
		Sources:  []string{"Associated Press", "Axios"},
		Language: "en",
	})

	if err != nil {
		panic(err)
	}

	for _, s := range articles.Articles {
		fmt.Printf("%+v\n\n", s)
	}
}
