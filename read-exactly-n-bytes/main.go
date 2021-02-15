package main

import (
	"io"
	"log"
	"os"
)

func main(){
	file, err := os.Open("test12.txt")
	if err !=nil {
		log.Fatal(err)
	}
	byteSlice := make([]byte, 2)
	writeByteSlice, err := io.ReadFull(file, byteSlice)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Number of bytes read: %d\n", writeByteSlice)
	log.Printf("Data read: %s\n", byteSlice)
}
