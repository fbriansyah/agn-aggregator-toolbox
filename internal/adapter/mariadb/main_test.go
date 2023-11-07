package mariadb

import (
	"database/sql"
	"log"
	"testing"

	"github.com/fbriansyah/agn-aggregator-toolbox/util"
	_ "github.com/go-sql-driver/mysql"
)

var (
	testQueries *Queries
	testDB      *sql.DB
	testAdapter DatabaseAdapter
)

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err := sql.Open("mysql", config.DBSource)
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}
	testQueries = New(testDB)
	testAdapter = NewDatabaseAdapter(testDB)
}
