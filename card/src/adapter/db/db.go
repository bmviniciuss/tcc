package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/lib/pq"
)

func ConnectDB() *pgxpool.Pool {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	dbname := os.Getenv("DATABASE_NAME")
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)
	dbpool, err := pgxpool.Connect(context.Background(), dsn)

	// log.Println("[db.ConnectDB] Connecting to database...")
	// db, err := sqlx.Connect("postgres", dsn)

	if err != nil {
		log.Fatalln("[db.ConnectDB] Error while connecting to database", err)
	}

	log.Println("[db.ConnectDB] Connected to database")

	return dbpool
}
