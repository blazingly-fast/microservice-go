package data

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
)

type Connection struct {
	conn *pgx.Conn
}

func NewConnection() *Connection {
	return &Connection{}
}

func Connect(sql string) {
	var dbconn *Connection = NewConnection()
	db, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Unable to establish connection: %v", err)
	}
	defer db.Close(context.Background())
	dbconn.conn = db
	tx, err := dbconn.conn.Begin(context.Background())
	if err != nil {
		log.Fatalf("conn.Begin failed: %v", err)
	}
	_, err = tx.Exec(context.Background(), sql)
	if err != nil {
		log.Fatal("tx.Exec failed: %v", err)
	}
	tx.Commit(context.Background())
}
