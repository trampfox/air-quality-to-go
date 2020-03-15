package scraper

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

type collyIpqaScraper struct{}

func CollyIPQAScraper() (IPQAScraper, error) {
	return &collyIpqaScraper{}, nil
}

// func (s *collyIpqaScraper) GetData() string {
// 	return ""
// }

func (s *collyIpqaScraper) GetStringData() string {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println(e.Attr("href"))
	})

	return ""
}
