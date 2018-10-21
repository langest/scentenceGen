package main

import (
	"flag"
	"log"
	"strings"
)

func main() {
	file := flag.String("-f", "", "The file where to load and save the probability map")

	flag.Parse()
	if 0 != strings.Compare(*file, "") {
		log.Fatal("TODO implement file management")
		return
	}

	log.Println(greedyMostProbableScentence(parseTrigramsFromStdin()))
}
