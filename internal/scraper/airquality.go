package scraper

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type RawDailyEntry []struct {
	NumColumns      int        `json:"numColumns"`
	Rows            [][]Row `json:"rows"`
	MarkFirstRow    bool       `json:"markFirstRow"`
	MarkFirstColumn bool       `json:"markFirstColumn"`
	HeaderLines     []struct {
		PrObject         string `json:"prObject"`
		NumberOfDecimals int    `json:"numberOfDecimals"`
		FontType         string `json:"fontType"`
		FontDimension    int    `json:"fontDimension"`
		Bold             bool   `json:"bold"`
		Italic           bool   `json:"italic"`
		Underlined       bool   `json:"underlined"`
		FixedWidth       bool   `json:"fixedWidth"`
		Center           bool   `json:"center"`
		NoWrap           bool   `json:"noWrap"`
		Highlighted      bool   `json:"highlighted"`
	} `json:"headerLines"`
	TailLines []interface{} `json:"tailLines"`
}

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

type Row string 

func (fr *Row) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err == nil {
		*fr = Row(s)
		return nil
	}

	var ov ObjectValue
	err = json.Unmarshal(b, &ov)
	if err == nil {
		*fr = Row(fmt.Sprintf("%.0f", ov.PrObject))
		return nil
	}

	var i int 
	err = json.Unmarshal(b, &i)
	if err == nil {
		*fr = Row(strconv.Itoa(i))
		return nil
	}
	
	return err
}
