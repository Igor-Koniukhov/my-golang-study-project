package main

import (
	"fmt"
	"log"
	"os"
)


type FooReade struct{}
type FooWrite struct{}

func (fooReade *FooReade) Reade(b []byte)(int, error){
	fmt.Print("in<")
	return os.Stdin.Read(b)
}

func (fooWrite *FooWrite) Write(b []byte)(int, error){
	fmt.Print("out>")
	return os.Stdout.Write(b)
}

func main(){
	var (
		read FooReade
		write FooWrite
	)
	input := make([]byte, 6500)
	s, err := read.Reade(input)
	if err !=nil {
		log.Fatal("Unable to read data")
	}
	fmt.Printf("Datf %d %T", s, s)

	s, err = write.Write(input)
	if err !=nil {
		log.Fatal("Unable to read data")
	}
	fmt.Printf("Data %d %T", s, s)
}
