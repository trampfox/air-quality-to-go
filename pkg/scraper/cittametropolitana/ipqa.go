package scraper

import (
	"regexp"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	ipqaUrl         = "http://www.cittametropolitana.torino.it/cms/ambiente/qualita-aria/dati-qualita-aria/ipqa"
	ipqaDescription = map[string]string{
		"iqa1": "ottima",
		"iqa2": "buona",
		"iqa3": "accettabile",
		"iqa4": "cattiva",
	}
	ipqaValue = map[string]string{
		"iqa1": "1",
		"iqa2": "2",
		"iqa3": "3",
		"iqa4": "4",
	}
)

type collyIpqaScraper struct{}

type IPQAData struct {
	Date                int64
	TodayValue          string
	TodayDescription    string
	TomorrowValue       string
	TomorrowDescription string
}

func CollyIPQAScraper() (IPQAScraper, error) {
	return &collyIpqaScraper{}, nil
}

func (s *collyIpqaScraper) GetStringData() IPQAData {
	ipqaData := IPQAData{
		Date: time.Now().Unix(),
	}
	c := colly.NewCollector()

	c.OnHTML("tr.valori", func(e *colly.HTMLElement) {
		re := regexp.MustCompile("iqa[1234]$")
		e.ForEach("td", func(index int, el *colly.HTMLElement) {
			// get iqa class used for the HTML element
			ipqaClass := re.FindString(el.Attr("class"))
			// the first column contains the today value
			if index == 0 {
				ipqaData.TodayValue = ipqaValue[ipqaClass]
				ipqaData.TodayDescription = ipqaDescription[ipqaClass]
				// whereas the second column contains the tomorrow value
			} else if index == 1 {
				ipqaData.TomorrowValue = ipqaValue[ipqaClass]
				ipqaData.TomorrowDescription = ipqaDescription[ipqaClass]
			}
		})
	})

	err := c.Visit(ipqaUrl)
	if err != nil {
		panic(err)
	}
	return ipqaData
}

// TODO implement check date function
// func checkDates(c *colly.Collector) bool {
// 	c.OnHTML("table.dati tbody tr", func(e *colly.HTMLElement) {
// 		e.ForEach("th", func(_ int, el *colly.HTMLElement) {
// 			fmt.Println(el.Text)
// 		})
// 	})

// 	return true
// }
