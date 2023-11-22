package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	if path := r.URL.Path; path != "/" {
		io.WriteString(w, "404 not found.")
		return 
	}
	io.WriteString(w, "Hello there.")
}

func getATR(w http.ResponseWriter, r *http.Request) {
	err, msg := parseQuery(r.URL.RawQuery)
	if (err == -1) {
		io.WriteString(w, msg + "\n")
		return 
	}
	
	period := err
	symbol := msg
	binanceURL := "https://testnet.binance.vision/api/v3/klines?symbol=" + symbol + "&interval=1d"
	response, error := http.Get(binanceURL)
	if (error != nil) {
		io.WriteString(w, "Error accessing binance api.")
		return 
	}

	resBody, _ := ioutil.ReadAll(response.Body)
	rows := strings.Split(string(resBody), "],[")
	
	var listOfStrings [][]string
	for _, row := range rows {
		column := strings.Split(row, ",")
		for i, col := range column {
			column[i] = strings.Replace(col, "\"", "", -1)
		}
		listOfStrings = append(listOfStrings, column)
	}

	var listOfOHLC [][]float64
	for _, str := range listOfStrings {
		open, _ := strconv.ParseFloat(str[1], 64)
		high, _ := strconv.ParseFloat(str[2], 64)
		low, _ := strconv.ParseFloat(str[3], 64)
		close, _ := strconv.ParseFloat(str[4], 64)
		listOfOHLC = append(listOfOHLC, []float64{open, high, low, close})
	}
	
	for i, j := 0, len(listOfOHLC)-1; i < j; i, j = i+1, j-1 {
		listOfOHLC[i], listOfOHLC[j] = listOfOHLC[j], listOfOHLC[i]
	}
	
	atrValue := calcATR(listOfOHLC, period)
	if atrValue == -1 {
		io.WriteString(w, "Invalid period. Data is insufficient.")
		return 
	}
	
	myMap := map[string]interface{}{
		"symbol": symbol,
		"period": period,
		"atr-value": atrValue,
	}

	jsonData, _ := json.Marshal(myMap)
	io.WriteString(w, string(jsonData))
}