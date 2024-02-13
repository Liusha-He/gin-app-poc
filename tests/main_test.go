package tests

import (
	"database/sql"
	"log"
	"os"
	"simple-bank/src/api"
	"simple-bank/src/dao"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

func NewTestServer(store dao.Store) (*api.Server, error) {
	config := api.Config{
		TokenSymmetricKey:   "12345678912345678912345678912345",
		AccessTokenDuration: time.Minute,
	}

	return api.NewServer(config, store)
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database")
	}
	testQueries = dao.New(testDB)

	os.Exit(m.Run())
}
