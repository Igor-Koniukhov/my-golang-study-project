package main

import (
	"log"
	"os"
)

func main(){
	err := os.Truncate("test0.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
