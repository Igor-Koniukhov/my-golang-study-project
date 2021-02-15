package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	fileinfo, err := os.Stat("testnew.txt")
	if err !=nil {
		if os.IsNotExist(err){
			log.Fatal("File does not exist")
		}
	}
	fmt.Println("File does exist. File information:")
	fmt.Println(fileinfo.Size())
}