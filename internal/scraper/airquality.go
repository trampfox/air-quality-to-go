package scraper

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

type RawDailyEntry []struct {
	NumColumns      int              `json:"numColumns"`
	Rows            [][]Row          `json:"rows"`
	MarkFirstRow    bool             `json:"markFirstRow"`
	MarkFirstColumn bool             `json:"markFirstColumn"`
	HeaderLines     []HeaderLine     `json:"headerLines"`
	TailLines       []RawTailLine `json:"tailLines"`
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

type HeaderLine struct {
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
}

type TailLine struct {
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
}

type Row string
type RawTailLine TailLine

func (fr *Row) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		*fr = Row(s)
		return nil
	}

	var ov ObjectValue
	if err := json.Unmarshal(b, &ov); err == nil {
		*fr = Row(fmt.Sprintf("%.1f", ov.PrObject))
		return nil
	}

	var hl HeaderLine
	if err := json.Unmarshal(b, &hl); err == nil {
		*fr = Row(hl.PrObject)
		return nil
	}

	var tl TailLine
	if err := json.Unmarshal(b, &tl); err == nil {
		*fr = Row(tl.PrObject)
		return nil
	}

	var i int
	if err := json.Unmarshal(b, &i); err == nil {
		*fr = Row(strconv.Itoa(i))
		return nil
	}

	log.Printf("unknown type, skip element %s", string(b))

	return nil
}

func (fr *RawTailLine) UnmarshalJSON(b []byte) error {

	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		*fr = RawTailLine{
			PrObject: s,
		}
		return nil
	}

	var tl TailLine
	if err := json.Unmarshal(b, &tl); err == nil {
		*fr = RawTailLine(tl)
		return nil
	}

	log.Printf("unknown type, skip element %s", string(b))

	return nil
}
