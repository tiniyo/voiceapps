package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
)

type UserIntent struct {
	Text          string          `json:"text"`
	Intent        Intent          `json:"intent"`
	Entities      []Entities      `json:"entities"`
	IntentRanking []IntentRanking `json:"intent_ranking"`
}
type Intent struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}
type Entities struct {
	Start      int    `json:"start"`
	End        int    `json:"end"`
	Text       string `json:"text"`
	Value      int    `json:"value"`
	Confidence int    `json:"confidence"`
	Entity     string `json:"entity"`
}
type IntentRanking struct {
	Name       string  `json:"name"`
	Confidence float64 `json:"confidence"`
}

func ProcessUserIntent(uIntent string) UserIntent {
	// userIntent is base64 encoded data.
	data, err := base64.StdEncoding.DecodeString(uIntent)
	if err != nil {
		log.Fatal("error:", err)
	}
	fmt.Println(string(data))
	// decode it to json
	var userIntent UserIntent
	json.Unmarshal(data, &userIntent)
	// parse json and get the intent
	fmt.Println(userIntent.Text, userIntent.Intent.Name)
	// process intent.
	return userIntent
}
