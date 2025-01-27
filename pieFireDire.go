package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

type MeatCount struct {
	Beef map[string]int `json:"beef"`
}

func countMeats(text string) map[string]int {
	counts := make(map[string]int)

	re := regexp.MustCompile(`[.,]+`)
	text = re.ReplaceAllString(text, " ")

	text = strings.Join(strings.Fields(text), " ")

	words := strings.Fields(strings.ToLower(text))

	for _, word := range words {
		if len(word) > 0 {
			counts[word]++
		}
	}

	return counts
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		http.Error(w, "Error fetching meat data", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading meat data", http.StatusInternalServerError)
		return
	}

	meatCounts := countMeats(string(body))

	result := MeatCount{
		Beef: meatCounts,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func runHttp() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
