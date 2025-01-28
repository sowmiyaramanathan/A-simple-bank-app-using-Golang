package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/api"
	db "github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:secret@localhost:5433/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
