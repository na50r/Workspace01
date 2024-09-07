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
