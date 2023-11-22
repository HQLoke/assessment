package main

import (
	"fmt"
	"net/http"
	"os"
)

var serverPort = 3456
var symbolList = []string{"BTCUSDT"}

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
