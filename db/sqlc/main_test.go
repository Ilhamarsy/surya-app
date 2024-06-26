package db

import (
	"database/sql"
	"log"
	"os"
	"projek-abal-abal/util"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("failed load config :", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed connect to db :", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
