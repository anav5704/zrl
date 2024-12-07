package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

    _ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
    godotenv.Load("./../.env")

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))

    if err != nil {
        fmt.Println("Error connecting to database: ", err)
        return nil, err
    }

    if err = db.Ping(); err != nil {
        fmt.Println("Error pinging database: ", err)
        return nil, err
    }

    _, err = db.Exec(schema)

    if err != nil {
        log.Println("Error creating schema: ", err)
        return nil, err
    }

    log.Println("Database initialized successfully")
    return db, nil
}
