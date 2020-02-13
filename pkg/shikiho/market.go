package shikiho

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Market struct {
	MarketCapitalization int     `json:"marketcap"`
	PBR                  float64 `json:"pbr"`
	PER                  float64 `json:"per"`
	NextPER              float64 `json:"next_per"`
	ROE                  float64 `json:"roe"`
	EquityRatio          float64 `json:"equity_ratio"`
}

func (market *Market) parseROE(doc *goquery.Selection) {
	doc.Each(func(i int, selection *goquery.Selection) {
		found := selection.Find("dt")

		if found.Text() == "ROE" {
			market.ROE = ParsePercentage(selection.Find("dd").Text())
		} else if found.Text() == "自己資本比率" {
			market.EquityRatio = ParsePercentage(selection.Find("dd").Text())
		}
	})
}

func (market *Market) parseStockInfo(doc *goquery.Selection) {
	//div.section div.block div.data div.stock
	doc.Each(func(i int, selection *goquery.Selection) {
		found := selection.Find("dt").First()

		if strings.Contains(found.Text(), "予想PER") {
			market.PER, market.NextPER = ParsePER(selection.Find("dd").First().Text())
		} else if found.Text() == "実績PBR" {
			market.PBR = ParsePBR(selection.Find("dd").First().Text())
		}
	})
}

func (market *Market) parseMarketCapitalization(doc *goquery.Selection) {
	market.MarketCapitalization = ParseJapanMoneyUnit(doc.Text())
}

func ParseMarket(doc *goquery.Selection) (*Market, error) {
	market := Market{}
	market.parseMarketCapitalization(doc.Find("div.overview div.sub div.stock dl.mc dd"))
	market.parseROE(doc.Find("div.sub div.block div.table dl"))
	market.parseStockInfo(doc.Find("div.section div.block div.data div.stock dl"))
	return &market, nil
}
