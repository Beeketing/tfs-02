package main

import (
	"../pkg"

	"../pkg/crawler"
)

func main() {
	filePath := "./output.log"
	logger, err := pkg.NewFileLogger(filePath)
	if err != nil {
		panic("cannot initialize logger")
	}
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "https://google.com/abcdefghijklmnopqrstuvwxyz"
	crawler.Crawl(url, sugar)
}
