package scraper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	airqualitytogo "github.com/trampfox/air-quality-to-go"
	"github.com/trampfox/air-quality-to-go/internal/scraper"
)

const (
	reportUrl = "http://www.cittametropolitana.torino.it/js/ariaweb/reports/rs-%s.json"
	pm10Name  = "PM10"
	pm25Name  = "PM2.5 - Beta"
)

type reportScraper struct {
	dateString string
}

type PollutionEntry struct {
	Date             string
	PollutantName    string
	ValueDescription string
	UnitOfMeasure    string
	DailyValues      []DailyValue
}

type DailyValue struct {
	Location      string
	Value         string
	ValueDateTime string
}

func ReportScraper(dateString string) (airqualitytogo.Scraper, error) {
	return &reportScraper{dateString: dateString}, nil
}

func (rs *reportScraper) GetData() string {
	dailyReportUrl := fmt.Sprintf(reportUrl, rs.dateString)
	log.Printf("Downloading report from %s", dailyReportUrl)

	response, err := http.Get(dailyReportUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	log.Println("Retrieve pollution entries...")
	var entries scraper.RawDailyEntry
	err = json.Unmarshal([]byte(contents), &entries)
	if err != nil {
		panic(err)
	}

	pEntries := rs.pollutionEntries(entries)

	b, err := json.Marshal(pEntries)
	if err != nil {
		panic(err)
	}

	return string(b)
}

func (rs *reportScraper) pollutionEntries(rawEntries scraper.RawDailyEntry) []PollutionEntry {
	var pEntries []PollutionEntry

	for _, rawEntry := range rawEntries {
		var dailyValues []DailyValue
		locationIndex := 0
		maxHourlyValueIndex := 1
		unitOfMeasureIndex := 1

		pollutantName := rawEntry.HeaderLines[0].PrObject
		if pollutantName != "" {
			// PM10 index exception
			if pollutantName == pm10Name {
				locationIndex = 1
				maxHourlyValueIndex = 2
				unitOfMeasureIndex = 2
			}

			// Retrieve daily values for the current pollutant
			for i := 2; i < len(rawEntry.Rows); i++ {
				dailyValue := DailyValue{
					Location:      string(rawEntry.Rows[i][locationIndex]),
					Value:         string(rawEntry.Rows[i][maxHourlyValueIndex]),
					ValueDateTime: rs.getValueDateTime(rawEntry.Rows[i], pollutantName),
				}
				dailyValues = append(dailyValues, dailyValue)
			}

			// Pollutant name is retrieved from the header line
			pEntry := PollutionEntry{
				PollutantName:    pollutantName,
				ValueDescription: rs.getValueDescription(rawEntry.Rows[0], pollutantName),
				UnitOfMeasure:    string(rawEntry.Rows[1][unitOfMeasureIndex]),
				DailyValues:      dailyValues,
			}
			pEntries = append(pEntries, pEntry)
		}
	}

	return pEntries
}

func (rs *reportScraper) getValueDateTime(rows []scraper.Row, pollutantName string) string {
	if pollutantName == pm10Name || pollutantName == pm25Name {
		return ""
	}

	return string(rows[2])
}

func (rs *reportScraper) getValueDescription(rows []scraper.Row, pollutantName string) string {
	descriptionIndex := 1
	if pollutantName == pm10Name {
		descriptionIndex = 2
	}

	return string(rows[descriptionIndex])
}
