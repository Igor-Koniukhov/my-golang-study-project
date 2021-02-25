package main

import (
	"io"
	"log"
	"net/http"
	"os"
)
var err = "error"
func main(){
	fl, err:= os.Create("devdungeon1.html")
	if err !=nil {
		log.Fatal(err)
	}
	defer fl.Close()
	url:= "https://www.devdungeon.com/content/working-files-go"
	conn, err:= http.Get(url)
	defer conn.Body.Close()
	cf, err := io.Copy(fl, conn.Body)
	checkError(err)
	log.Printf("Downloaded %d byte file.\n", cf)

}

func checkError(err error){
	if err !=nil {
		log.Fatal(err)
	}
}