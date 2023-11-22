package main

import (
	"strconv"
	"strings"
)

func parseQuery(query string) (int, string) {
	params := strings.Split(query, "&")
	var period int = 14
	var symbol string

	if len(params) <= 0 || len(params) >= 3 {
		return -1, "Wrong number of parameters."
	}

	if (len(params) == 2) {
		temp := strings.Split(params[1], "=")
		if len(temp) != 2 || temp[0] != "period" || len(temp[1]) == 0 {
			return -1, "Unknown parameter."
		}
		if isNumeric(temp[1]) == false {
			return -1, "Period is not integer value."
		}
		period, _ = strconv.Atoi(temp[1])
	}

	temp := strings.Split(params[0], "=")
	if len(temp) != 2 || temp[0] != "symbol" || len(temp[1]) == 0 {
		return -1, "Unknown parameter."
	}
	if strings.Contains(strings.Join(symbolList, ","), temp[1]) == false {
		return -1, "Symbol not found."
	}
	symbol = temp[1]
	return period, symbol
}
