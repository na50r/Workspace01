package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const (
	set1 = "exampleSet.json"
)

var dataPath = "d:/Personal/Workspace01/server/data"

func createFilePath(folderPath string, file string) (filePath string) {
	filePath = filepath.Join(folderPath, file)
	return
}

func fetchWordCombis(folderpath string, file string) (wordCombis []WordCombination, err error) {
	filepath := createFilePath(folderpath, file)
	jsonFile, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var combiSet CombiSet
	err = json.Unmarshal(jsonFile, &combiSet)
	if err != nil {
		return nil, err
	}

	return combiSet.Data, nil
}

func InitSet1() {
	wordCombis, err := fetchWordCombis(dataPath, set1)
	if err != nil {
		panic(err)
	}

	store, err := NewPostgresStore()
	if err != nil {
		panic(err)
	}

	for _, combi := range wordCombis {
		err := store.CreateWordCombi(&combi)
		if err != nil {
			panic(err)
		}
	}
}
