package shikiho

import (
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
	"strings"
)

type PerformanceOfYear struct {
	// e.g. 連19.3予
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
	Year   int    `json:"year"`
	Month  int    `json:"month"`
	// 売上
	// 営業利益
	// 経常利益
	// 純利益
	// 1株益
	// 1株配
	Sales            int     `json:"sales"`
	OperatingIncome  int     `json:"operating_income"`
	OrdinaryProfit   int     `json:"ordinary_profit"`
	NetIncome        int     `json:"net_income"`
	ProfitPerStock   float64 `json:"profit_per_stock"`
	DividendPerStock float64 `json:"dividend_per_stock"`
}

func ParsePerformance(doc *goquery.Selection) ([]PerformanceOfYear, error) {
	table := ParseTable(doc)

	regex := regexp.MustCompile("^\\s*(連|単|◎|◇|□|中|会)?(\\d{2})\\.(\\d{1,2})[\\*※]?(予|変)*[\\*※]?\\s*$")
	performances := []PerformanceOfYear{}

	for i := 1; i < len(table); i++ {
		if len(table[i]) != 7 {
			return nil, errors.New("table cols is not 7")
		}

		values := table[i]

		if !regex.MatchString(values[0]) {
			// 連19.4〜9
			if strings.Contains(values[0], "〜") {
				continue
			}

			for j := 0; j < len(values); j++ {
				fmt.Fprintf(os.Stderr, "Values: %s\n", values[j])
			}
			return nil, errors.New(fmt.Sprintf("cannot parse table year: %s", values[0]))
		}

		perform := PerformanceOfYear{}

		group := regex.FindStringSubmatch(values[0])
		// year
		perform.Prefix = group[1]
		perform.Year = ParseInt(group[2])
		perform.Month = ParseInt(group[3])
		perform.Suffix = group[4]

		if perform.Prefix == "会" || perform.Prefix == "中" {
			continue
		}

		perform.Sales = ParseInt(values[1])
		perform.OperatingIncome = ParseInt(values[2])
		perform.OrdinaryProfit = ParseInt(values[3])
		perform.NetIncome = ParseInt(values[4])
		perform.ProfitPerStock = ParseFloat(values[5])
		perform.DividendPerStock = ParseFloat(values[6])

		performances = append(performances, perform)
	}

	return performances, nil
}

func MergePerformances(performancesOld, performancesUpdate []PerformanceOfYear) []PerformanceOfYear {
	results := performancesOld

	for i := 0; i < len(performancesUpdate); i++ {
		update := performancesUpdate[i]
		found := false

		for j := 0; j < len(performancesOld); j++ {
			old := performancesOld[j]

			if old.Year == update.Year && old.Month == update.Month && old.Prefix == update.Prefix {
				results[j] = update
				found = true
				break
			}
		}

		if !found {
			results = append(results, update)
		}
	}

	return results
}
