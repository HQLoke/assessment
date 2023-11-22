package main

import (
	// "fmt"
	"math"
	"unicode"
)

type dailyPrice struct {
	openTime					int			
	openPrice 					float64
	highPrice 					float64
	lowPrice  					float64
	closePrice					float64
	volume						float64
	closeTime					int
	quoteAssetVolume			float64
	numOfTrades					int
	takerBuyBaseAssetVolume		float64
	takerBuyQuoteAssetVolume	float64
	ignore						int
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
** The first argument, p, is a array of daily prices.
** The second argument, period is required for atr calculation.
**
** acronyms:
** hml = today's high minus low
** hmcp = absolute value of today's high minus yesterday's closing price
** lmcp = absolute value of today's low minus yesterday's closing price
 */
func calcATR(p [][]float64, period int) float64 {
	if period >= len(p) {
		return -1
	}

	sum := float64(0)
	for i := 0; i < period; i += 1 {
		hml := p[i][1] - p[i][2]
		hmcp := math.Abs(p[i][1] - p[i+1][3])
		lmcp := math.Abs(p[i][2] - p[i+1][3])

		max := findLargest(hml, hmcp, lmcp)
		sum += max
	}
	return sum / float64(period)
}
