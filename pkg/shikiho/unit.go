package shikiho

import (
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
