package main

import (
	"bytes"
	"log"
	"math/rand"
	"strings"
	"time"
	"errors"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	bigramProbabilityMap map[string]map[string]int = make(map[string]map[string]int)
)

func addBigramToProbabilityMap(s1, s2 string) {
	log.Println("Adding bigram probability:", s1, s2)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	if bigramProbabilityMap[s1] == nil {
		bigramProbabilityMap[s1] = make(map[string]int)
	}
	bigramProbabilityMap[s1][s2] += 1
}

func guessNextWord(str string) (string, error) {
	log.Println("Guessing next word")
	sum := 0
	for _, occurance := range bigramProbabilityMap[str] {
		sum += occurance
	}
	log.Println("A")
	r := rand.Intn(sum)
	log.Println("r:", r)

	for word, occurance := range bigramProbabilityMap[str] {
		log.Println(word,":", occurance)
		r -= occurance
		if r < 0 {
			return word, nil
		}
	}
	return "", errors.New("Error, code should not reach this point")
}

func greedyMostProbableScentence() string {
	log.Println("Greedy probable scentence search")
	var buffer bytes.Buffer
	for word, _ := guessNextWord(""); word != ""; word, _ = guessNextWord(word) { //TODO handle error
		buffer.WriteString(word)
		buffer.WriteString(" ")
	}
	return buffer.String()
}
