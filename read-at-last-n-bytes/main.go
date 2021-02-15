package main

import (
	"io"
	"log"
	"os"
)
func main (){
	file, err := os.Open("test12.txt")
	if err !=nil {
		log.Fatal(err)
	}
	bs := make([]byte, 12)
	minb := 8
	readAtLast, err := io.ReadAtLeast(file, bs, minb)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", readAtLast)
	log.Printf("Data read: %s\n", bs)
}