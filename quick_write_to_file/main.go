package main

import (
	"io/ioutil"
	"log"
)

func main() {
	err := ioutil.WriteFile("test12.txt", []byte("\nHi there!"), 0666)
	if err !=nil {
		log.Fatal(err)

	}
}
