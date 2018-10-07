package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		previousWord := ""
		for _, word := range strings.Split(string(line), " ") {
			addBigramToProbabilityMap(previousWord, word)
			previousWord = word
		}
		addBigramToProbabilityMap(previousWord, "") //Add probability for ending on word
		if isPrefix {
			log.Println("Scentence too long")
			//TODO Save the day
		}
		line, isPrefix, err = reader.ReadLine()
	}
	log.Println(greedyMostProbableScentence())
}
