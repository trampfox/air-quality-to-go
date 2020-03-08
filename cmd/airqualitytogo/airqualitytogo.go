package main

import (
	"fmt"

	"github.com/trampfox/air-quality-to-go/pkg/scraper"
)

func main() {
	// TODO get date from argument
	rs, err := scraper.ReportScraper("20200305")
	if err != nil {
		panic(err)
	}

	data := rs.GetData()
	fmt.Println(data)
}
