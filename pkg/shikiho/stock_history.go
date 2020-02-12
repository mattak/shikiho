package shikiho

import "github.com/PuerkitoBio/goquery"

type StockHistory struct {
	Histories []StockHistoryOfYear `json:"histories"`
}

type StockHistoryOfYear struct {
	Year    int    `json:"year"`
	Month   int    `json:"month"`
	Content string `json:"content"`
	Size    string `json:"size"`
}

func (s *StockHistory) ParseHistories(doc *goquery.Selection) {
	values := ParseTable(doc)
	s.Histories = []StockHistoryOfYear{}

	for i := 0; i < len(values); i++ {
		fields := values[i]
		year, month, err := ParseYearMonth(fields[0])
		if err != nil {
			continue
		}

		s.Histories = append(s.Histories, StockHistoryOfYear{
			Year:    year,
			Month:   month,
			Content: fields[1],
			Size:    fields[2],
		})
	}
}
