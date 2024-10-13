package database

import (
    "context"
    "github.com/jackc/pgx/v5/pgxpool"
    "log"
)

var dbPool *pgxpool.Pool

func ConnectDB(connStr string) {
    var err error
    dbPool, err = pgxpool.New(context.Background(), connStr)
    if err != nil {
        log.Fatalf("unable to connect to database: %v", err)
    }
    log.Println("Connected to the database")
}

func GetDB() *pgxpool.Pool {
    return dbPool
}
