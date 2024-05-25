package main

import (
	"database/sql"
	"log"
	"simple-bank/src/api"
	"simple-bank/src/dao"
	"time"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	address  = "localhost:8083"
)

// @title     simple bank API
// @version         1.0
// @description     A Golang and gin API template
func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to database!", err)
	}

	store := dao.NewStore(conn)

	// todo - update the tokensymmetrickey for production (dev for now)
	config := api.Config{
		TokenSymmetricKey:   "12345678912345678912345678912345",
		AccessTokenDuration: 15 * time.Minute,
	}

	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(address)
	if err != nil {
		log.Fatal("cannot start server!", err)
	}
}
