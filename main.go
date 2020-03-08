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
	UnitOfMeasure string
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

	pEntries := pollutionEntries(entries)

	b, err := json.Marshal(pEntries)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}

func pollutionEntries(rawEntries scraper.RawDailyEntry) []PollutionEntry {
	var pEntries []PollutionEntry
	
	for _, rawEntry := range rawEntries {
		var dailyValues []DailyValue
		locationIndex := 0
		maxHourlyValueIndex := 1
		unitOfMeasureIndex := 1

		pollutantName := rawEntry.HeaderLines[0].PrObject
		if pollutantName != "" {
			// PM10 index exception
			if pollutantName == "PM10" {
				locationIndex = 1
				maxHourlyValueIndex = 2
				unitOfMeasureIndex = 2
			}
			if len(rawEntry.Rows[1]) > 0 {
				fmt.Printf("%v\n", rawEntry.Rows[1])
			}
			// Retrieve daily values for the current pollutant
			for i := 2; i < len(rawEntry.Rows); i++ {
				dailyValue := DailyValue{
					Location:       string(rawEntry.Rows[i][locationIndex]),
					MaxHourlyValue: string(rawEntry.Rows[i][maxHourlyValueIndex]),
					// MaxHourlyValueDateTime: string(rawEntry.Rows[i][2]),
				}
				dailyValues = append(dailyValues, dailyValue)
			}

			// Pollutant name is retrieved from the header line
			pEntry := PollutionEntry{
				PollutantName: pollutantName,
				UnitOfMeasure: string(rawEntry.Rows[1][unitOfMeasureIndex]),
				DailyValues:   dailyValues,
			}
			pEntries = append(pEntries, pEntry)
		}
	}

	return pEntries
}
