package shikiho

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"os"
)

type Journal struct {
	Code         string              `json:"code"`
	Market       Market              `json:"market"`
	Performances []PerformanceOfYear `json:"performances"`
}

func ParseJournal(filepath string) (*Journal, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}

	market, err := ParseMarket(doc.Selection)
	if err != nil {
		return nil, err
	}

	performances, err := ParsePerformance(doc.Find("div.performance div.matrix"))
	if err != nil {
		return nil, err
	}

	performancesUpdate, err := ParsePerformance(doc.Find("div.update div.matrix"))
	if err != nil {
		return nil, err
	}

	performances = MergePerformances(performances, performancesUpdate)

	code := doc.Find("div.main div.title div.code").Text()

	return &Journal{
		Code:         code,
		Market:       *market,
		Performances: performances,
	}, nil
}

func (journal Journal) ToJson() string {
	result, err := json.Marshal(journal)
	if err != nil {
		panic(err)
	}
	return string(result)
}
