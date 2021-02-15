package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main (){
	file, err := os.Open("test12.txt")
	if err !=nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Printf("Data as hex: %x\n", data)
	fmt.Printf("Data as string: %s\n", data)
	fmt.Println("Number of bytes read:", len(data))
}
