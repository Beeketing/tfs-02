package crawler

import (
	"io/ioutil"
	"net/http"
	"sync"

	"time"

	"go.uber.org/zap"
)

func Crawl(wg *sync.WaitGroup, url string, sugar *zap.SugaredLogger) {
	// for synching between goroutines
	if wg != nil {
		defer wg.Done()
	}
	//  make sure that all thing happens here not cause application panic
	defer func() {
		if r := recover(); r != nil {
			sugar.Warn("Recovered from Crawl function")
		}
	}()

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
		return
	}
	// close body after reading content
	defer resp.Body.Close()
	// read all content using ioutil.ReadAll function
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		sugar.Errorf("Cannot read body due to error, ", err)
		return
	}
	sugar.Infof("Downloaded content length %v", len(b))
}
