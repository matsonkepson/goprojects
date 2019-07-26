package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatalf("Usage: %s <port>", os.Args[0])
	}
	if _, err := strconv.Atoi(os.Args[1]); err != nil {
		log.Fatalf("Invalid port: %s (%s)\n", os.Args[1], err)
	}
	/* 	if first, err := strconv.ParseInt(os.Args[1], 10); err != nil {
		fmt.Println("This prots is restricted\nTry something above 1024", os.Args[1], err)
	} */

	/* 	if os.Args[1] <= init(1024) {
		fmt.Println("This prots is restricted\nTry something above 1024")
		os.Exit(1)
	} */

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "i am ready, %q", html.EscapeString(req.URL.Path))
		println("--->", os.Args[1], req.URL.String())
	})
	http.ListenAndServe(":"+os.Args[1], nil)
}
