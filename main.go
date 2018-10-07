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
		for _, sentence := range tokenizer.Tokenize(string(line)) {
			trimmed := strings.Trim(sentence.Text, " ")
			words := strings.Split(trimmed, " ")
			if len(words) < 3 /* || len(words) > 10*/ {
				continue
			}
			word0 := words[0]
			word1 := words[1]
			word2 := words[2]
			addTrigramToProbabilityMap("", "", word0)
			addTrigramToProbabilityMap("", word0, word1)
			for i := 0; i < len(words)-2; i++ {
				word0 = words[i]
				word1 = words[i+1]
				word2 = words[i+2]
				addTrigramToProbabilityMap(word0, word1, word2)
			}
			addTrigramToProbabilityMap(word1, word2, "") //Add probability for ending on word
		}
		if isPrefix {
			//TODO Save the day
		}
		line, isPrefix, err = reader.ReadLine()
	}
	log.Println(greedyMostProbableScentence())
}
