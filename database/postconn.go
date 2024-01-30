package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresqlConnection struct {
	Url string
}

func NewDatabaseConnection(databaseUrl string) *PostgresqlConnection {
	return &PostgresqlConnection{
		Url: databaseUrl,
	}
}

func (p *PostgresqlConnection) Connect() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv(p.Url))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	var greeting string

	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)
}
