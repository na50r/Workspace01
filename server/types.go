package main

type WordCountReqest struct {
	Word string `json:"word"`
}

type WordCountResponse struct {
	Count int `json:"count"`
}

type WordCount struct {
	Word  string
	Count int
}

type WordCombination struct {
	Word1  string `json:"word1"`
	Word2  string `json:"word2"`
	Result string `json:"res"`
}

type CombiSet struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Data        []WordCombination `json:"data"`
}
