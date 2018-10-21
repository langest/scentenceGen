package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	loadPath := flag.String("f", "", "The file where to load and save the probability map")
	savePath := flag.String("s", "", "The file where to save the probability map, will terminate the program after saving the map")
	flag.Parse()

	var trigrams *TrigramProbabilityMap
		trigrams = parseTrigramsFromStdin()
	if 0 != strings.Compare(*loadPath, "") {
		//trigrams = NewTrigramProbabilityMap()
		trigrams.LoadFromJsonFile(*loadPath)
	} else {
		trigrams = parseTrigramsFromStdin()
	}

	if 0 != strings.Compare(*savePath, "") {
		trigrams.SaveToJsonFile(*savePath)
		return
	}

	log.Println(greedyMostProbableScentence(trigrams))
}
