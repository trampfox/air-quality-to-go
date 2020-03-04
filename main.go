package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/trampfox/air-quality-to-go/internal/scraper"
)

type ObjectValue struct {
	PrObject         float64 `json:"prObject"`
	NumberOfDecimals int     `json:"numberOfDecimals"`
	Bold             bool    `json:"bold"`
	Italic           bool    `json:"italic"`
	Underlined       bool    `json:"underlined"`
	FixedWidth       bool    `json:"fixedWidth"`
	Center           bool    `json:"center"`
	NoWrap           bool    `json:"noWrap"`
	Highlighted      bool    `json:"highlighted"`
}

type PollutionEntry struct {
	Date          string
	PollutantName string
	DailyValues   []DailyValue
}

type DailyValue struct {
	Location               string
	MaxHourlyValue         string
	MaxHourlyValueDateTime string
}

func main() {
	response, err := http.Get("http://www.cittametropolitana.torino.it/js/ariaweb/reports/rs-20200224.json")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var entries scraper.RawDailyEntry
	err = json.Unmarshal([]byte(contents), &entries)
	if err != nil {
		panic(err)
	}

	pollutantNames(entries)
	fmt.Printf("%v", measuringStations(entries))

	// for _, entry := range entries {
	// 	fmt.Printf("%v\n", entry.HeaderLines)
	// }

	// fmt.Printf("%v\n", entries)
	// fmt.Printf("%s\n", string(contents))
}

func pollutantNames(rawEntries scraper.RawDailyEntry) []PollutionEntry {
	var pEntries []PollutionEntry
	for _, rawEntry := range rawEntries {
		if pName := rawEntry.HeaderLines[0].PrObject; pName != "" {
			pEntry := PollutionEntry{
				PollutantName: rawEntry.HeaderLines[0].PrObject,
			}
			pEntries = append(pEntries, pEntry)
		}
	}
	fmt.Printf("%v\n", pEntries)
	return pEntries
}

func measuringStations(rawEntries scraper.RawDailyEntry) []DailyValue {
	var dailyValues []DailyValue
	for _, rawEntry := range rawEntries {
		// fmt.Printf("%v\n", rawEntry.Rows)
		for i := 2; i < len(rawEntry.Rows); i++ {
			dailyValue := DailyValue{
				Location:       string(rawEntry.Rows[i][0]),
				MaxHourlyValue: string(rawEntry.Rows[i][1]),
			}
			dailyValues = append(dailyValues, dailyValue)
		}
	}

	return dailyValues
}
