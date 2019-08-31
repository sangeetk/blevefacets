package main

import (
	"encoding/json"
	"log"

	"github.com/blevesearch/bleve"
)

type Blog struct {
	Title  string "json:title"
	Author string "json:author"
	Body   string "json:body"
}

func main() {
	mapping := bleve.NewIndexMapping()
	index, err := bleve.NewMemOnly(mapping)
	if err != nil {
		panic(err)
	}

	blogs := []Blog{
		{Title: "Hello World", Author: "Sangeet Kumar", Body: "Hello world post"},
		{Title: "Hello World Again", Author: "Mayan Sangeet", Body: "Hello world post again"},
	}

	// Index blogs
	for _, b := range blogs {
		index.Index(b.Title, b)
	}

	// Simple text search
	query := bleve.NewQueryStringQuery("hello")
	searchRequest := bleve.NewSearchRequest(query)
	searchRequest.Fields = []string{"*"}
	searchResult, err := index.Search(searchRequest)
	if err != nil || searchResult.Total == 0 {
		log.Println("Not found")
		return
	}

	for _, hit := range searchResult.Hits {
		jsonstr, _ := json.Marshal(hit.Fields)
		log.Println(hit.ID, string(jsonstr))
	}

}
