package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)
func main(){
	file, err:= os.Open("testnew.txt")
	if err !=nil {
		log.Fatal(err)

	}
	file.Close()

	file, err = os.OpenFile("testnew.txt", os.O_APPEND, 0666)
	if err !=nil {
		log.Fatal(err)

	}

	fileStat, err := os.Stat("testnew.txt")
	if err !=nil {
		log.Fatal(err)
	}
	fmt.Println(fileStat.Mode())
	msg := strings.NewReader("\n hello this new text ")

	io.Copy(file, msg )
	file.Close()

}