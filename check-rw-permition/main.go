package main

import (
	"fmt"
	"log"
	"os"
)
func main(){
	file, err := os.OpenFile("testnew.txt", os.O_WRONLY, 0666)
	if err !=nil {
		if os.IsPermission(err){
			log.Println("Error: Write permission denied")
		}
	}
	file.Close()
	file, err = os.OpenFile("testnew.txt", os.O_RDONLY, 0666)
	if err !=nil {
		if os.IsPermission(err){
			log.Println("Error: Read permission denied")
		}
	}
	file.Close()
	stat, err := os.Stat("testnew.txt")
	if err !=nil {
		log.Println("Error: Permission denied")
	}
	fmt.Println(stat.Mode())

}