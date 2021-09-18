package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Println("start")
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
        fmt.Fprint(w, "Hello world!")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
