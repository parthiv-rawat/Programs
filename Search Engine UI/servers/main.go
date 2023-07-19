package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/rs/cors"
)

type SearchResult struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func searchResults(query string, results []SearchResult) any {
	var searchResults []SearchResult

	for _, result := range results {
		if strings.Contains(strings.ToLower(result.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(result.Description), strings.ToLower(query)) {
			searchResults = append(searchResults, result)
		}
	}

	if searchResults == nil {
		return "This search result does not exist."
	}
	return searchResults
}

func main() {

	results := []SearchResult{
		{
			Title:       "Harward University",
			Description: "This is the official website of Harward University",
			URL:         "https://www.harvard.edu/",
		},
		{
			Title:       "Stanford University",
			Description: "This is the official website of Stanford University",
			URL:         "https://www.stanford.edu/",
		},
		{
			Title:       "University of Oxford",
			Description: "This is the official website of Oxford University",
			URL:         "https://www.ox.ac.uk/",
		},
		{
			Title:       "Massachusetts Institute of Technology",
			Description: "This is the official website of MIT",
			URL:         "https://www.mit.edu/",
		},
		{
			Title:       "University of Cambridge",
			Description: "This is the official website of Cambridge University",
			URL:         "https://www.cam.ac.uk/",
		},
		{
			Title:       "National University of Singapore",
			Description: "This is the official website of NUS",
			URL:         "https://nus.edu.sg/",
		},
		{
			Title:       "Princeton University",
			Description: "This is the official website of Princeton University",
			URL:         "https://www.princeton.edu/",
		},
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println(r.URL.Query())
		query := r.URL.Query().Get("query")

		// fmt.Println(query)
		searchResults := searchResults(query, results)

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(searchResults)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	handler := cors.Default().Handler(mux)

	log.Fatal(http.ListenAndServe(":8081", handler))
}
