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
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Println("[db.ConnectDB] Connecting to database...")
	dbpool, err := pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		log.Fatalln("[db.ConnectDB] Error while connecting to database", err)
	}

	log.Println("[db.ConnectDB] Connected to database")

	return dbpool
}
