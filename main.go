package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/data"
)

func main() {
	b, _ := data.Asset("data/english.json")
	training, _ := sentences.LoadTraining(b)
	tokenizer := sentences.NewSentenceTokenizer(training)

	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		log.Println("Parsing line:\n", line)
		previousWord := ""
		for _, sentence := range tokenizer.Tokenize(string(line)) {
			log.Println("Parsing sentence:\n", sentence)
			for _, word := range strings.Split(sentence.Text, " ") {
				addBigramToProbabilityMap(previousWord, word)
				previousWord = word
			}
			addBigramToProbabilityMap(previousWord, "") //Add probability for ending on word
			if isPrefix {
				log.Println("Scentence too long")
				//TODO Save the day
			}
		}
		line, isPrefix, err = reader.ReadLine()
	}
	log.Println(greedyMostProbableScentence())
}
