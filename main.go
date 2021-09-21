package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

var DB *sql.DB

// DBへの接続
func connectDB() (*sql.DB, error) {
    dbUser := os.Getenv("DB_USER")
    dbPass := os.Getenv("DB_PASS")
    dbName := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")

    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?multiStatements=true", dbUser, dbPass, dbHost, dbName))

    if err != nil {
        return nil, fmt.Errorf("err %s", err)
    }

    return db, nil
}

// マイグレーションの実行
func migration() error {
    log.Println("start migration")
    driver, err := mysql.WithInstance(DB, &mysql.Config{})

    if err != nil {
        return fmt.Errorf("err %s", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "mysql",
        driver,
    )

    if err != nil {
        return fmt.Errorf("err %s", err)
    }

    err = m.Up()

    if err != nil {
        return fmt.Errorf("err %s", err)
    }

    return nil
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello world! Welcome Go App!")
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
        http.HandleFunc("/health_check", HealthCheck)
        log.Fatal(http.ListenAndServe(":8080", nil))
    case "migrate":
        err := migration()
        if err != nil {
            log.Fatal(err)
        }
    }
}
