package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"website/favicon"
	"website/robots"
	"website/tarpit"
)

func main() {
	port := flag.Uint("port", 3000, "port to listen on")
	flag.Parse()

	http.HandleFunc("/favicon.ico", favicon.Handler)
	http.HandleFunc("/robots.txt", robots.Handler)
	http.HandleFunc("/", tarpit.Handler)

	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
