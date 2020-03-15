package scraper

// Scraper is an interface for different type of scraper objects
type Scraper interface {
	GetStringData() string
	GetData() []PollutionEntry
}

type IPQAScraper interface {
	GetStringData() string
}