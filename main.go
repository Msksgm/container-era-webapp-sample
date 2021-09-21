package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world! Welcome Go App!!")
}

func HealthCheck(w http.ResponseWriter, r *http.Request)  {
    fmt.Fprint(w, "ok")
}

func main() {
    if len(os.Args) < 2 {
        fmt.Println("help")
        os.Exit(0)
    }

	log.Println("start")

    db, err := connectDB()

    if err != nil {
        log.Fatal(err)
    }

    DB = db
    defer DB.Close()

    switch os.Args[1] {
    case "start":
        http.HandleFunc("/", RootHandler)
        http.HandleFunc("/todo", PostHandler)
        http.HandleFunc("/health_check", HealthCheck)
        log.Fatal(http.ListenAndServe(":8080", nil))
    case "migrate":
        err := migration()
        if err != nil {
            log.Fatal(err)
        }
    }
}
