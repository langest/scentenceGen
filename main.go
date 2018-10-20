package main

import (
	"flag"
	"log"
)

func main() {
	addKorpus := flag.Bool("add-korpus", false, "If true, parses stdin and saves to json file") //TODO better mesesage

	flag.Parse()
	if *addKorpus {
		log.Println("TODO add korpus")
		return
	}

	log.Println(greedyMostProbableScentence(parseTrigramsFromStdin()))
}
