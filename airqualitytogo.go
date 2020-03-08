package airqualitytogo

// Scraper is an interface for different type of downloader objects 
type Scraper interface {
	GetData() string
}