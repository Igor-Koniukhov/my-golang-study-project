package main

import (
	"io"
	"log"
	"os"
)

func main(){
	originalFile, err := os.Open("test12.txt")
	if err !=nil {
		log.Fatal(err)
	}
	defer originalFile.Close()
	newFile, err := os.Create("test13.txt")
	if err !=nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	bytesWriter, err := io.Copy(newFile, originalFile)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWriter)
	err = newFile.Sync()
	if err !=nil {
		log.Fatal(err)
	}
}
