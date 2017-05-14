package main

import (
	"fmt"
	"github.com/badoux/goscraper"
	"net/http"
	"encoding/json"
)

type PREVIEW struct {
	title string
	description string
}

func getPreview(w http.ResponseWriter, r *http.Request) {
	s, err := goscraper.Scrape(r.URL.Query().Get("url"), 5)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.NewEncoder(w).Encode(s)
}

func main() {
	http.HandleFunc("/", getPreview)
	http.ListenAndServe(":18088", nil)
}
