package main

import (
	"bytes"
	"errors"
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

var (
	trigramProbabilityMap map[string]map[string]map[string]int = make(map[string]map[string]map[string]int)
)

func addTrigramToProbabilityMap(s1, s2, s3 string) {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	s3 = strings.ToLower(s3)
	if trigramProbabilityMap[s1] == nil {
		trigramProbabilityMap[s1] = make(map[string]map[string]int)
	}
	if trigramProbabilityMap[s1][s2] == nil {
		trigramProbabilityMap[s1][s2] = make(map[string]int)
	}
	trigramProbabilityMap[s1][s2][s3] += 1
}

func guessNextWord(s1, s2 string) (string, error) {
	sum := 0
	for _, occurance := range trigramProbabilityMap[s1][s2] {
		sum += occurance
	}
	r := rand.Intn(sum + 1)

	for word, occurance := range trigramProbabilityMap[s1][s2] {
		r -= occurance
		if r <= 0 {
			return word, nil
		}
	}
	return "", errors.New("Error, code should not reach this point")
}

func greedyMostProbableScentence() string {
	var buffer bytes.Buffer
	word0 := ""
	word1 := ""
	word2, _ := guessNextWord("", "")
	for ; word2 != ""; word2, _ = guessNextWord(word0, word1) { //TODO handle error
		buffer.WriteString(word2)
		buffer.WriteString(" ")
		word0 = word1
		word1 = word2
	}
	buffer.WriteString(word2)
	return buffer.String()
}
