package shikiho

import "github.com/PuerkitoBio/goquery"

type Market struct {
	Code                 string  `json:code`
	MarketCapitalization int     `json:marketcap`
	Pbr                  float64 `json:pbr`
	Per                  float64 `json:per`
}

func ParseMarket(doc *goquery.Selection) (*Market, error) {
	mc := doc.Find("div.overview div.sub div.stock dl.mc dd").Text()

	market := Market{}
	market.MarketCapitalization = ParseJapanMoneyUnit(mc)

	return &market, nil
}
