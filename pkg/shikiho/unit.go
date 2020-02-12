package shikiho

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func ParseInt(value string) int {
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, "―", "0", -1)
	value = strings.Replace(value, "‥", "0", -1)

	r, err := strconv.ParseInt(value, 10, 64)
	if err != nil {
		panic(err)
	}
	return int(r)
}

func ParseFloat(value string) float64 {
	value = strings.Replace(value, "記", "", -1)
	value = strings.Replace(value, "特", "", -1)
	value = strings.Replace(value, "*", "", -1)
	value = strings.Replace(value, ",", "", -1)
	value = strings.Replace(value, "‥", "0", -1)
	value = strings.Replace(value, "--", "0", -1)

	if i := strings.Index(value, "〜"); i != -1 {
		value = value[0:i]
	}

	r, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot parse: \"%s\"\n", value)
		panic(err)
	}

	return r
}

func ParseJapanMoneyUnit(raw string) int {
	raw = strings.Replace(raw, ",", "", -1)

	// e.g. 1,234億円
	regex := regexp.MustCompile("\\s*([\\d\\.\\,]+)\\s*(\\S+)\\s*$")

	if !regex.MatchString(raw) {
		if raw == "--" {
			return 0
		}
		panic(fmt.Sprint("parse japan money unit error: %s", raw))
	}

	groups := regex.FindStringSubmatch(raw)
	value := ParseInt(groups[1])
	unit := groups[2]

	if unit == "億円" {
		return value
	}

	panic(fmt.Sprintf("unknown unit: %s", unit))
}

func ParsePercentage(raw string) float64 {
	regex := regexp.MustCompile("\\s*(\\-?[\\d\\.\\,]+)\\s*%\\s*")

	if !regex.MatchString(raw) {
		fmt.Fprintf(os.Stderr, "parse percentage failed: %s\n", raw)
		return 0
	}

	group := regex.FindStringSubmatch(raw)
	value := ParseFloat(group[1])

	return value
}

func ParsePER(raw string) (float64, float64) {
	regex1 := regexp.MustCompile("^\\s*(\\-?[\\d\\.\\,]+|\\-\\-)\\s*\\((\\-?[\\d\\.\\,]+|\\-\\-)\\)\\s*")

	if !regex1.MatchString(raw) {
		fmt.Fprintf(os.Stderr, "parse per failed: %s\n", raw)
		return 0, 0
	}

	group := regex1.FindStringSubmatch(raw)
	value1 := ParseFloat(group[1])
	value2 := ParseFloat(group[2])
	return value1, value2
}

func ParsePBR(raw string) float64 {
	regex := regexp.MustCompile("^\\s*(\\-?[\\d\\.\\,]+)\\s*倍\\s*")

	if !regex.MatchString(raw) {
		fmt.Fprintf(os.Stderr, "parse pbr failed: %s\n", raw)
		return 0
	}

	group := regex.FindStringSubmatch(raw)
	value := ParseFloat(group[1])
	return value
}

func ParseYearMonth(value string) (int, int, error) {
	regex := regexp.MustCompile("^\\s*(\\d+)\\.(\\d+)\\s*$")

	group := regex.FindStringSubmatch(value)
	if len(group) > 0 {
		year, _ := strconv.Atoi(group[1])
		month, _ := strconv.Atoi(group[2])
		return year, month, nil
	}

	return 0, 0, errors.New(fmt.Sprintf("parse failed year, month: %s", value))
}
