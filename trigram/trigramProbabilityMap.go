package trigram

import (
	"strings"
	"encoding/json"
	"log"
)

type TrigramProbabilityMap struct {
	occurrences map[string]map[string]map[string]int
}

func NewTrigramProbabilityMap() *TrigramProbabilityMap {
	return &TrigramProbabilityMap{occurrences: make(map[string]map[string]map[string]int)}
}

func (this TrigramProbabilityMap) LoadFromJsonFile(path string) {
	jsonBlob := readBytesFromFile(path)
	var tmp TrigramProbabilityMap
	err := json.Unmarshal(jsonBlob, &tmp.occurrences)
	if err != nil {
		log.Println("Error: Failed to unmarshal json")
		return
	}
	this.MergeWith(&tmp)
}

func (this TrigramProbabilityMap) SaveToJsonFile(path string) {
	jsonBlob, err := json.Marshal(this.occurrences)
	if err != nil {
		log.Println("Error: Failed to convert trigrams to json object,\n", err)
		return
	}
	writeBytesToFile(jsonBlob, path)
}

func (this TrigramProbabilityMap) Get(s0, s1 string) map[string]int {
	if this.occurrences[s0] != nil && this.occurrences[s0][s1] != nil {
		return this.occurrences[s0][s1]
	}
	return nil
}

func (this TrigramProbabilityMap) AddTrigramOccurrence(s0, s1, s2 string, val int) {
	s0 = strings.ToLower(s0)
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)
	if this.occurrences[s0] == nil {
		this.occurrences[s0] = make(map[string]map[string]int)
	}
	if this.occurrences[s0][s1] == nil {
		this.occurrences[s0][s1] = make(map[string]int)
	}
	this.occurrences[s0][s1][s2] += val
}

func (this TrigramProbabilityMap) MergeWith(other *TrigramProbabilityMap) {
	for str0, map1 := range other.occurrences {
		for str1, map2 := range map1 {
			for str2, val := range map2 {
				this.AddTrigramOccurrence(str0, str1, str2, val)
			}
		}
	}
}
