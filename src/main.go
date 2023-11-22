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
	fmt.Println("Server is listening on port", serverPort, "...")
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
