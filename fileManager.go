package main

import (
	"encoding/json"
	"log"
)

func saveMapToFile(trigrams TrigramProbabilityMap, filename string) {
	r, rr := json.Marshal(trigrams)
	log.Println(string(r))
	log.Println(rr)
	log.Fatal("TODO save to file,", filename)
}

func loadMapFromfile(filename string) *TrigramProbabilityMap {
	log.Fatal("TODO implement")
	return NewTrigramProbabilityMap()
}
