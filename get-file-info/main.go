package main

import (
	"fmt"
	"log"
	"os"
)
var (
	fileInfo os.FileInfo
	err error
)

func main(){
	fileInfo, err := os.Stat("test0.txt")
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println("File name:", fileInfo.Name())
	fmt.Println("File size", fileInfo.Size())
	fmt.Println("Permitions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory:", fileInfo.IsDir())
	fmt.Printf("System interface type: %T/n", fileInfo.Sys())
	fmt.Printf("System info: %+v\n\n", fileInfo.Sys())
}