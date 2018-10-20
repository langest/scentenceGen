package main

import (
	"log"

	"bytes"
	"encoding/json"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func greedyMostProbableScentence(trigrams TrigramProbabilityMap) string {
	var buffer bytes.Buffer
	word0 := ""
	word1 := ""
	word2, _ := guessNextWord("", "", trigrams)
	var err error
	for ; word2 != ""; word2, err = guessNextWord(word0, word1, trigrams) {
		if err != nil {
			log.Fatal(err) //TODO handle better?
		}
		buffer.WriteString(word2)
		buffer.WriteString(" ")
		word0 = word1
		word1 = word2
	}
	buffer.WriteString(word2)
	return buffer.String()
}

func guessNextWord(s1, s2 string, trigrams TrigramProbabilityMap) (string, error) {
	sum := 0
	for _, occurance := range trigrams[s1][s2] {
		sum += occurance
	}
	r := rand.Intn(sum + 1)

	for word, occurance := range trigrams[s1][s2] {
		r -= occurance
		if r <= 0 {
			return word, nil
		}
	}
	return "", errors.New("Error, code should not reach this point")
}
