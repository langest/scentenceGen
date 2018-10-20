package main

import (
	"strings"
)

type TrigramProbabilityMap map[string]map[string]map[string]int

func addTrigramToProbabilityMap(s1, s2, s3 string, trigrams TrigramProbabilityMap) {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	s3 = strings.ToLower(s3)
	if trigrams[s1] == nil {
		trigrams[s1] = make(map[string]map[string]int)
	}
	if trigrams[s1][s2] == nil {
		trigrams[s1][s2] = make(map[string]int)
	}
	trigrams[s1][s2][s3] += 1
}

