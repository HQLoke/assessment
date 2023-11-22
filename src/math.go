package main

import (
	// "fmt"
	"math"
	"unicode"
)

type dailyPrice struct {
	openPrice  float64
	highPrice  float64
	lowPrice   float64
	closePrice float64
}

func isNumeric(s string) bool {
	for _, c := range s {
		if unicode.IsDigit(c) == false {
			return false
		}
	}
	return true
}

func findLargest(nums ...float64) float64 {
	largest := float64(0)

	for _, num := range nums {
		if num > largest {
			largest = num
		}
	}
	return largest
}

/*
** The first argument, p, is a array of daily prices. It differs from a normal array
** because the most recent data is at the beginning of the array.
**
** acronyms:
** hml = today's high minus low
** hmcp = absolute value of today's high minus yesterday's closing price
** lmcp = absolute value of today's low minus yesterday's closing price
 */
func calcATR(p []dailyPrice, period int) float64 {
	if period >= len(p) {
		return -1
	}

	sum := float64(0)
	for i := 0; i < period; i += 1 {
		hml := p[i].highPrice - p[i].lowPrice
		hmcp := math.Abs(p[i].highPrice - p[i+1].closePrice)
		lmcp := math.Abs(p[i].lowPrice - p[i+1].closePrice)

		max := findLargest(hml, hmcp, lmcp)
		sum += max
	}
	return sum / float64(period)
}
