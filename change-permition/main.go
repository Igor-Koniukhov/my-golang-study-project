package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main(){
	err := os.Chmod("test-hasing.txt", 0777)
	if err !=nil {
		log.Println(err)
	}
	//change ownership
	err = os.Chown("test-hasing.txt", os.Getuid(), os.Getgid())
	if err !=nil {
		log.Println(err)
	}
	//change timestamps
	twoDaysFromNow := time.Now().Add(48 * time.Hour)
	lastAccessTime := twoDaysFromNow
	lastModifyTime := twoDaysFromNow
	err = os.Chtimes("test-hasing.txt", lastAccessTime, lastModifyTime)
	if err !=nil {
		log.Println(err)
	}
	stat, err := os.Stat("test-hasing.txt")
	if err !=nil {
		log.Println(err)
	}
	fmt.Println(stat.Mode(), stat.ModTime())
}