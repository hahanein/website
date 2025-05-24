package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := flag.String("port", "", "port to listen on (required)")
	flag.Parse()

	if *port == "" {
		log.Fatal("port flag is required")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		if _, err := fmt.Fprintln(w, "Hello World"); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Printf("listening on :%s", *port)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
