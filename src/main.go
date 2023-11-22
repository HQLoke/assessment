package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	// "regexp"
)

var serverPort = 3456
var symbolList = []string{"BTCUSDT"}

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
	
	// period := err
	// symbol := msg
	binanceURL := "https://testnet.binance.vision/api/v3/klines?symbol=BTCUSDT&interval=1d"
	response, error := http.Get(binanceURL)
	if (error != nil) {
		io.WriteString(w, "Error accessing binance api.")
		os.Exit(1)
	}
	resBody, _ := ioutil.ReadAll(response.Body)
	output := fmt.Sprintf("%s", resBody)
	io.WriteString(w, output + "HAHAHAHAHA")
}

func main() {
	mux	:= http.NewServeMux()
	mux.HandleFunc("/", getRoot)
	mux.HandleFunc("/api/indicator/atr", getATR)

	addr	:= fmt.Sprintf(":%d", serverPort)
	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
