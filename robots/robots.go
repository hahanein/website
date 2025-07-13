package robots

import (
	"fmt"
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	response := `
User-agent: *
Disallow: /
	`

	if _, err := fmt.Fprint(w, response); err != nil {
		log.Printf("handling request failed: %v", err)
	}
}
