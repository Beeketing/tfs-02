package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	fmt.Println("This is a fmt message")
	log.Println("This is a log message")

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "https://google.com/abcdefghijklmnopqrstuvwxyz"
	sugar.Infof("Crawling url %s", url)
	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		sugar.Errorw("failed to fetch URL",
			// Structured context as loosely typed key-value pairs.
			"url", url,
			"attempt", 1,
			"backoff", time.Second,
			"error", err,
		)
	}
}
