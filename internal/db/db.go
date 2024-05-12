package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/garnet-geo/garnet-userapi/internal/env"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var conn *pgxpool.Pool

func InitConnection() {
	var err error
	conn, err = pgxpool.New(context.Background(), env.GetDatabaseUrl())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	log.Default().Println("Connected to the database")
}

func CloseConnection() {
	conn.Close()

	log.Default().Println("Closed connection to the database")
}

func ExecuteQuery(query string, args ...any) (pgx.Rows, error) {
	return conn.Query(context.Background(), query, args...)
}

func ExecuteQueryRow(query string, args ...any) pgx.Row {
	return conn.QueryRow(context.Background(), query, args...)
}
