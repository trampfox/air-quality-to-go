package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/prometheus/common/log"
	"github.com/trampfox/air-quality-to-go/internal/scraper"
)

func main() {
	response, err := http.Get("http://www.cittametropolitana.torino.it/js/ariaweb/reports/rs-20200224.json")
	if err != nil {
		log.Error(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error(err)
	}

	var entries scraper.RawDailyEntry
	err = json.Unmarshal([]byte(contents), &entries)
	if err != nil {
		log.Error(err)
	}

	fmt.Printf("%v\n", entries[0])

	// for _, entry := range entries {
	// 	fmt.Printf("%v\n", entry.HeaderLines)
	// }

	// fmt.Printf("%v\n", entries)
	// fmt.Printf("%s\n", string(contents))
}
