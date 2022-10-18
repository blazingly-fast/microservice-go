package data

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

func Connection() *pgx.Conn {
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	// defer db.Close(context.Background())
	return db
}
