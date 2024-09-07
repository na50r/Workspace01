package main

import (
	"fmt"

	"github.com/na50r/server/sse"
)

func checkWordWithDB(broker *sse.Broker, word string) (count int, err error) {

	store, err := NewPostgresStore()
	if err != nil {
		return
	}

	wc, found, err := store.GetWordCount(word)
	if err != nil {
		return
	}

	if !found {
		if word[0] == 'a' {
			broker.BroadcastEvent("Word with A detected")
		}

		if err = store.CreateWordCount(&WordCount{Word: word, Count: len(word)}); err != nil {
			return
		}
		return len(word), nil
	} else {
		fmt.Println("CACHED")
		count = wc.Count
		return
	}
}
