package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"website/favicon"
)

func main() {
	port := flag.Uint("port", 3000, "port to listen on")
	flag.Parse()

	http.HandleFunc("/favicon.ico", favicon.Handler)
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		file, err := os.Open("/proc/uptime")
		if err != nil {
			log.Printf("Error opening /proc/uptime: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		if _, err := file.WriteTo(w); err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	log.Printf("listening on :%d", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
