package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"math/rand"
	"io/ioutil"
)

const (
	maxLuckyNumber = 60
)

type Fortune struct {
	Text         string `json:"text"`
	LuckyNumbers []int  `json:"luckyNumbers"`
}

func main() {
	fmt.Println("Fortune service started")

	rand.Seed(42)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fortune := Fortune{getFortune(), generateLuckyNumbers()}

		json.NewEncoder(w).Encode(fortune)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func generateLuckyNumbers() []int {
	return []int{rand.Intn(maxLuckyNumber), rand.Intn(maxLuckyNumber), rand.Intn(maxLuckyNumber), rand.Intn(maxLuckyNumber), rand.Intn(maxLuckyNumber)}
}

func getFortune() string {
	raw, err := ioutil.ReadFile("fortunes.json")
	if err != nil {
		panic("Error reading fortunes.json file")
	}

	var fortunes []string
	json.Unmarshal(raw, &fortunes)

	return fortunes[rand.Intn(len(fortunes))]
}
