package main

import (
    "fmt"
    "net/http"
    "encoding/json"
    "github.com/badoux/goscraper"
)

func getPreview(w http.ResponseWriter, r *http.Request) {
    url := r.URL.Query().Get("url")

    s, err := goscraper.Scrape(url, 5)
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
