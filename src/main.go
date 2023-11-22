package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	// "regexp"
)

const serverPort = 3456

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}

func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func	calcATR(w http.ResponseWriter, r *http.Request) {

}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)

	addr	:= fmt.Sprintf(":%d", serverPort)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}