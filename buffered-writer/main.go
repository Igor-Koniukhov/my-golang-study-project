package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	bufferedWriter := bufio.NewWriter(file)
	bytesWriten, err := bufferedWriter.Write([]byte{34,56,89, 57, 89})
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWriten)
	bytesWriten, err = bufferedWriter.WriteString(" Let it be string.... more string")
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Bytes written: %d\n", bytesWriten)
	//how much stored in buffered file
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)
	//how much buffer is available
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Bytes available: %d\n", bytesAvailable)
	//write memory buffer to disk
	bufferedWriter.Flush()
	//change to a new writer
	bufferedWriter.Reset(bufferedWriter)
	//see how much buffer is available
	bytesAvailable = bufferedWriter.Available()
	if err !=nil {
		log.Fatal(err)
	}
	log.Printf("Available buffer: %d\n", bytesAvailable)

	bufferedWriter = bufio.NewWriterSize(bufferedWriter, 10000,)
	bytesAvailable = bufferedWriter.Available()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Available byffer: %d\n", bytesAvailable)


}