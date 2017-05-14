package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/badoux/goscraper"
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
	http.ListenAndServe(":9090", nil)
}
