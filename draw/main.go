package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file, err := os.Create("draw.txt")
	if err !=nil {
		log.Fatal(err)
	}
	defer file.Close()
	n, err := file.Write([]byte("hello draw!"))
	if err !=nil {
		log.Fatal(err)
	}
	c, err := file.WriteString("\nanother hello!")
	if err !=nil {
		log.Fatal(err)
	}
	b, err := file.WriteAt([]byte("this is offset"), 27)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("File info: bytes %d, %d, %d", n, c, b)
	fmt.Println(n,c,b, "some info")
	flsendler, err := os.Open("original_also.txt")
	if err !=nil {
		log.Fatal(err)
	}
	flstat, err := flsendler.Stat()
	if err !=nil {
		log.Fatal(err)
	}
	bs:= make([]byte, flstat.Size())
	_,err = flsendler.Read(bs)
	if err !=nil {
		log.Fatal(err)
	}
	g, err := file.Write(bs)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Data of g %d,", g)



}
