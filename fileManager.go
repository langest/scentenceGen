package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

func saveMapToFile(trigrams TrigramProbabilityMap, filename string) (err error) {
	var json string
	json, err = json.Marshal(trigrams)
	if err != nil {
		log.Println("Failed to convert trigrams to json object:", err)
		return
	}
	err = writeStringToFile(string(json), "myTestFile.json")
	if err != nil {
		log.Println("Failed to write json string to file:", string(json))
	}
	return
}

func loadMapFromfile(filename string) *TrigramProbabilityMap {
	log.Fatal("TODO implement")
	return NewTrigramProbabilityMap()
}

func readStringFromFile(path string) (string) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}
	return string(file)
}

func writeStringToFile(content string, path string) (err error) {
	var file *os.File
	if file, err = os.Create(path); err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	return
}
