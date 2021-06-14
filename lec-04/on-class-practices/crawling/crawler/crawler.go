package crawler

type Crawler interface {
	Crawl(string) (string, error)
}
