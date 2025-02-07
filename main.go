package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/api"
	db "github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/db/sqlc"
	"github.com/sowmiyaramanathan/A-simple-bank-app-using-Golang/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to database : ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server")
	}
}
