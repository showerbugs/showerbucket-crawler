package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "./crawl"
)

func getPreview(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")

    s, err := crawl.Scrape(url, 5)
    if err != nil {
        return
    }

    json.NewEncoder(w).Encode(s)
}

func main() {
    http.HandleFunc("/", getPreview)
    http.ListenAndServe(":9090", nil)
}
