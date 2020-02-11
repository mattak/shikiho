package shikiho

import (
	"fmt"
	"math"
	"strings"
)

type PerformanceSorter struct {
	Code         string
	Performances []PerformanceOfYear
}

func (p *PerformanceSorter) Len() int {
	return len(p.Performances)
}

func (p *PerformanceSorter) Less(i, j int) bool {
	diffYear := p.Performances[i].Year - p.Performances[j].Year
	if diffYear < 0 {
		return true
	}
	diffMonth := p.Performances[i].Month - p.Performances[j].Month
	if diffYear == 0 && diffMonth < 0 {
		return true
	}
	return false
}

func (p *PerformanceSorter) Swap(i, j int) {
	tmp := p.Performances[i]
	p.Performances[j] = tmp
	p.Performances[i] = p.Performances[j]
}

func (p PerformanceSorter) LastReal() PerformanceOfYear {
	for i := len(p.Performances) - 1; i >= 0; i-- {
		if p.Performances[i].Suffix != "予" {
			return p.Performances[i]
		}
	}
	return p.Performances[0]
}

func (p PerformanceSorter) LastAny() PerformanceOfYear {
	return p.Performances[len(p.Performances)-1]
}

func (p PerformanceSorter) ToTsv() string {
	first := p.Performances[0]
	last := p.LastReal()
	spanDiff := (last.Year-first.Year)*12 + (last.Month - first.Month)

	if spanDiff < 0 {
		panic(fmt.Sprintf("span diff is illegal: %s %d.%d -> %d.%d", p.Code, first.Year, first.Month, last.Year, last.Month))
	}

	if !first.IsSequentialRange(last) {
		panic(fmt.Sprintf("not matching prefix: %s %s%d.%d%s %s%d.%d%s",
			p.Code,
			first.Prefix, first.Year, first.Month, first.Suffix,
			last.Prefix, last.Year, last.Month, last.Suffix,
		))
	}

	var salesGrowth float64
	var operatingIncomeGrowth float64
	var netIncomeGrowth float64
	yearDiff := float64(spanDiff / 12.0)

	if spanDiff > 0 {
		salesGrowth = math.Pow(float64(last.Sales)/float64(first.Sales), 1.0/yearDiff)
		operatingIncomeGrowth = math.Pow(float64(last.OperatingIncome)/float64(first.OperatingIncome), 1.0/yearDiff)
		netIncomeGrowth = math.Pow(float64(last.NetIncome)/float64(first.NetIncome), 1.0/yearDiff)
	} else {
		salesGrowth = 1
		operatingIncomeGrowth = 1
		netIncomeGrowth = 1
	}

	return strings.Join([]string{
		p.Code,
		fmt.Sprintf("%d.%d", first.Year, first.Month),
		fmt.Sprintf("%d.%d", last.Year, last.Month),
		fmt.Sprintf("%d", spanDiff),
		fmt.Sprintf("%f", salesGrowth),
		fmt.Sprintf("%f", operatingIncomeGrowth),
		fmt.Sprintf("%f", netIncomeGrowth),
	}, "\t")
}
