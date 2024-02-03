package tests

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/src/dao"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var (
	testQueries *dao.Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}
	testQueries = dao.New(testDB)

	os.Exit(m.Run())
}
