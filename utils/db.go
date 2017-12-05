package utils

import (
    "database/sql"
    _ "github.com/lib/pq"
    "gopkg.in/doug-martin/goqu.v4"
    _ "gopkg.in/doug-martin/goqu.v4/adapters/postgres"
    "log"
    "os"
)

func SetupDB() *goqu.Database {
    log.Println(os.Getenv("DATABASE_URL"))
    db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatal(err)
        return &goqu.Database{}
    }

    return goqu.New("postgres", db)
}