package main

import (
	"log"
	"os"
)

func main(){
	file, err := os.OpenFile("testfile.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666)
	if err !=nil {
		log.Fatal(err)
	}
	defer file.Close()

	msg := []byte("massage to new file.... a lot of letters / how long can be that string/ Can it be more the 10000000000 b")
	bytesWritten, err := file.Write(msg)
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Wrote %d bytes.\n", bytesWritten)

}