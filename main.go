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
    http.HandleFunc("/health_check", func (w http.ResponseWriter, r *http.Request)  {
        fmt.Fprint(w, "ok")
    })
    log.Fatal(http.ListenAndServe(":8080", nil))
}
