package main

import (
	"bytes"
	"strings"
)

var (
	bigramProbabilityMap map[string]map[string]int = make(map[string]map[string]int)
)

func addBigramToProbabilityMap(s1, s2 string) {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
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
