package book

import (
	"encoding/json"
	"fmt"
)

// Book defines a book's attributes and basic operation and interaction with Elasticsearch
type Book struct {
	ID            string `json:"id,omitempty"`
	ISBN          string `json:"isbn"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	NumberOfPage  int64  `json:"number_of_page"`
	PublishedDate string `json:"published_date"`
	Tags          string `json:"tags"`
	Brief         string `json:"brief"`
}

// String returns object's string representation
func (book *Book) String() string {
	if book == nil {
		return ""
	}
	b, err := json.Marshal(book)
	if err != nil {
		fmt.Println("Cannot convert to json: ", err)
		return ""
	}
	return string(b)
}
