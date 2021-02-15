package main

import (
	"io/ioutil"
	"log"
)

func main(){
	//file read to byte slice
	data, err := ioutil.ReadFile("test13.txt")
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Data read: %s\n", data)
}