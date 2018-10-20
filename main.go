package main

import (
	"flag"
	"log"
)

func main() {
	file := flag.String("-f", "", "The file where to load and save the probability map")

	flag.Parse()
	if file != "" {
		log.Fatal("TODO implement file management")
		return
	}

	log.Println(greedyMostProbableScentence(parseTrigramsFromStdin()))
}
