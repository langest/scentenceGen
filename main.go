package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"strings"
)

var (
	bigramProbabilityMap map[string]map[string]int = make(map[string]map[string]int)
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

func addBigramToProbabilityMap(s1, s2 string) {
	if bigramProbabilityMap[s1] == nil {
		bigramProbabilityMap[s1] = make(map[string]int)
	}
	bigramProbabilityMap[s1][s2] += 1
}

func guessNextWord(str string) (bestGuess string, probability float64) {
	bestGuess = ""
	max := -1
	sum := 0
	for word, occurances := range bigramProbabilityMap[str] {
		if occurances > max {
			max = occurances
			bestGuess = word
		}
		sum += occurances
	}
	probability = float64(max) / float64(sum)
	return
}

func greedyMostProbableScentence() string {
	var buffer bytes.Buffer
	for word, _ := guessNextWord(""); word != ""; word, _ = guessNextWord(word) {
		buffer.WriteString(word)
		buffer.WriteString(" ")
	}
	return buffer.String()
}
