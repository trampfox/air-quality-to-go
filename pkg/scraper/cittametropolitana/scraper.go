package scraper

import "github.com/trampfox/air-quality-to-go/internal/scraper"

// Scraper is an interface for different type of scraper objects
type Scraper interface {
	GetStringData() string
	GetData() []PollutionEntry
}

type IPQAScraper interface {
	GetStringData() scraper.IPQAData
}
