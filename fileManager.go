package main

import (
	"io/ioutil"
	"log"
)

func readBytesFromFile(path string) (bytes []byte) {
	var err error
	bytes, err = ioutil.ReadFile(path)
	if err != nil {
		log.Println("Error:", err)
	}
	return
}

func writeBytesToFile(bytes []byte, path string) {
	var err error
	err = ioutil.WriteFile(path, bytes, 0644)
	if err != nil {
		log.Println("Error:", err)
	}
}
