package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"website/favicon"
)

func main() {
	port := flag.Uint("port", 3000, "port to listen on")
	flag.Parse()

	http.HandleFunc("/favicon.ico", favicon.Handler)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello World"); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
