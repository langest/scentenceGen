package main

import (
	"bufio"
	"os"
	"log"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, isPrefix, err := reader.ReadLine()
	for err == nil {
		log.Println(string(line))
		if isPrefix {
			log.Println()
		}
		line, isPrefix, err = reader.ReadLine()
	}
}
