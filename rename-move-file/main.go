package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main(){
	originalPath := "test0.txt"
	newPath := "testnew.txt"
	err := os.Rename(originalPath, newPath)
	if err !=nil {
		log.Fatal(err)
	}
	msg:=[]byte( "New content for new documentation")

	ioutil.WriteFile("testnew.txt", msg, 6)

}