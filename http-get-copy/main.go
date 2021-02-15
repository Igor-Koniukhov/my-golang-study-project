package main

import (

	"io"
	"log"
	"net/http"
	"os"
)
var err = "error"
func main(){
	newFile, err:= os.Create("rodim.html")
	if err !=nil {
		log.Fatal(err)
	}
	defer newFile.Close()
	url:= "https://rodim.com.ua/"
	conn, err:= http.Get(url)
	defer conn.Body.Close()
	cf, err := io.Copy(newFile, conn.Body)
	if err !=nil {
		log.Fatal(err)

	}
	log.Printf("Downloaded %d byte file.\n", cf)




}
