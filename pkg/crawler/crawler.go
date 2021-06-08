package crawler

type Crawler interface {
	RulesFound() bool
	Crawl() []byte
	Save() bool
}
