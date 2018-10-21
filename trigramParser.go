package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"

	"gopkg.in/neurosnap/sentences.v1"
	"gopkg.in/neurosnap/sentences.v1/data"
)

func parseTrigramsFromStdin() (ret *TrigramProbabilityMap) {
	ret = NewTrigramProbabilityMap()
	b, _ := data.Asset("data/english.json")
	training, _ := sentences.LoadTraining(b)
	tokenizer := sentences.NewSentenceTokenizer(training)

	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		for _, sentence := range tokenizer.Tokenize(string(line)) {
			trimmed := trimUnwanted(strings.Trim(sentence.Text, " "))
			words := strings.Split(trimmed, " ")
			if len(words) < 3 {
				continue
			}
			word0 := words[0]
			word1 := words[1]
			word2 := words[2]
			ret.AddTrigramOccurrence("", "", word0, 1)
			ret.AddTrigramOccurrence("", word0, word1, 1)
			for i := 0; i < len(words)-2; i++ {
				word0 = words[i]
				word1 = words[i+1]
				word2 = words[i+2]
				ret.AddTrigramOccurrence(word0, word1, word2, 1)
			}
			ret.AddTrigramOccurrence(word1, word2, "", 1) //Add probability for ending on word
		}
		if isPrefix {
			//TODO Save the day
			log.Fatal("NOOOO")
		}
		line, isPrefix, err = reader.ReadLine()
	}
	return
}

func trimUnwanted(str string) string {
	reg, err := regexp.Compile("[^a-zA-Z0-9i_\\-,;:\\.\\såäöÅÄÖ]+")
	if err != nil {
		log.Fatal(err)
	}
	return reg.ReplaceAllString(str, "")
}
