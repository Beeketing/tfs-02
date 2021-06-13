package main

import (
	"sync"
	"time"

	"../pkg"
	"../pkg/crawler"
	"go.uber.org/zap"
)

/*
  In this application, we will use go-routine to demonstrate the process of crawling webpages concurrently
*/
func main() {
	urls := []string{
		"https://techcrunch.com/2021/06/12/apples-ipados-15-breaks-the-app-barrier/",
		"https://techcrunch.com/2021/06/11/the-air-taxi-market-prepares-to-take-flight/",
		"https://techcrunch.com/2021/06/12/how-many-opinions-does-it-take-to-hit-the-100m-arr-club/",
		"https://techcrunch.com/2021/06/12/7-new-security-features-apple-quietly-announced-at-wwdc/",
		"https://techcrunch.com/2021/06/11/the-rise-of-robotaxis-in-china/",
		"https://techcrunch.com/2021/06/11/facebook-buys-game-studio-bigbox-vr/",
		"https://techcrunch.com/2021/06/11/extra-crunch-roundup-eu-insurtech-30-years-of-crossing-the-chasm-embedded-finances-endgame/",
	}
	// initilize logger
	logger, _ := pkg.NewFileLogger("")
	sugar := logger.Sugar()
	startTime := time.Now()
	downloadWithoutGoroutine(urls, sugar)
	period := time.Since(startTime)
	sugar.Info("Download without goroutine time: ", period)

	startTime = time.Now()
	downloadWithGoroutine(urls, sugar)
	period = time.Since(startTime)
	sugar.Info("download with goroutine time: ", period)
}

// Download from provided urls without using goroutine
func downloadWithoutGoroutine(urls []string, sugar *zap.SugaredLogger) {
	for _, url := range urls {
		crawler.Crawl(nil, url, sugar)
	}
}

func downloadWithGoroutine(urls []string, sugar *zap.SugaredLogger) {
	wg := sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go crawler.Crawl(&wg, url, sugar)
	}
	wg.Wait()
}
