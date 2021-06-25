package book

import (
	"errors"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v7"
)

type ESClient struct {
	*elastic.Client
}

// NewESClient creates new es client based on url connection
func NewESClient(url string) (*ESClient, error) {
	if len(url) == 0 {
		return nil, errors.New("empty url connection")
	}
	client, err := elastic.NewClient(elastic.SetURL(url),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false))

	fmt.Println("ES initialized...")

	return &ESClient{client}, err
}
