package ent

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func Migrate() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dns := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbDatabase,
	) + "parseTime=true&collation=utf8mb4_bin"

	client, err := Open("mysql", dns)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

}
