package shikiho

import "github.com/PuerkitoBio/goquery"

func ParseTableLine(elements *goquery.Selection) []string {
	values := []string{}

	elements.Each(func(i int, selection *goquery.Selection) {
		values = append(values, selection.Text())
	})

	return values
}

func ParseTable(table *goquery.Selection) [][]string {
	result := [][]string{}

	table.Find("thead tr").Each(func(i int, s *goquery.Selection) {
		result = append(result, ParseTableLine(s.Find("th")))
	})

	table.Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		result = append(result, ParseTableLine(s.Find("td")))
	})

	return result
}
