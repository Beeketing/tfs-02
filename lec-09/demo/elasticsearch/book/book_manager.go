package book

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"time"

	elastic "gopkg.in/olivere/elastic.v7"
)

const (
	indexName = "books"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type BookManager struct {
	esClient *ESClient
}

func NewBookManager(es *ESClient) *BookManager {
	return &BookManager{esClient: es}
}

func (bm *BookManager) SearchBooks(title string) []*Book {
	ctx := context.Background()

	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return nil
	}
	// build query to search for title
	query := elastic.NewSearchSource()
	query.Query(elastic.NewMatchQuery("title", title))

	// get search's service
	searchService := bm.esClient.
		Search().
		Index(indexName).
		SearchSource(query)

	// perform search query
	searchResult, err := searchService.Do(ctx)
	if err != nil {
		fmt.Println("Cannot perform search with ES", err)
		return nil
	}
	// get result
	var books []*Book

	for _, hit := range searchResult.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			fmt.Println("Get data error: ", err)
			continue
		}
		fmt.Println(&book)
		books = append(books, &book)
	}

	return books
}

func (bm *BookManager) AddBook(book *Book) error {
	ctx := context.Background()
	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	// fill ID if it's empty
	if len(book.ID) == 0 {
		fmt.Println("emtpy ID")
		book.ID = rndStringRunes(10)
	}

	_, err := bm.esClient.Index().
		Index(indexName).
		BodyJson(book.String()).
		Id(book.ID).
		Do(ctx)

	// call to flush data to disk for search. if no call --> need to wait for 5s to search since inserted
	bm.esClient.Refresh(indexName).Do(ctx)

	return err
}

func (bm *BookManager) DeleteBook(book *Book) error {
	ctx := context.Background()
	if bm.esClient == nil {
		fmt.Println("Nil es client")
		return errors.New("nil es client")
	}

	res, err := bm.esClient.Delete().
		Index(indexName).
		Id(book.ID).
		Do(ctx)

	if res.Shards.Successful > 0 {
		fmt.Println("Document deleted from from index: ", book.ID)
	}
	return err
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func rndStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
