package main

import (
	"fmt"
	"log"
	"os"

)
var (
	newFile *os.File
	err error
)
type Data struct{
	x, y int
}
func sumOfNum(d *Data) int{

	return d.x + d.y
}
func mult(d *Data) int{
	return d.x*d.y
}

func main(){
	newFile, err:= os.Create("test0.txt")
	if err !=nil {
		log.Fatal(err)

	}
	log.Println(newFile)
	newFile.Close()

	d:= Data{3,8}
	fmt.Println(sumOfNum(&d), mult(&d))




}