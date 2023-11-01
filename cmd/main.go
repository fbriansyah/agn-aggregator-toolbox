package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/fiber"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/adapter/mariadb"
	"github.com/fbriansyah/agn-aggregator-toolbox/internal/application"
)

func main() {
	db, err := sql.Open("mysql", "root:secret@tcp(localhost:3306)/db_aggregator_dev?parseTime=true")
	if err != nil {
		log.Fatalf("cannot connect to database: %v", err)
	}

	dbAdapter := mariadb.NewDatabaseAdapter(db)

	service := application.NewService(dbAdapter)
	// template := template.NewTemplateAdapter()

	// httpAdapter := chi.NewChiAdapter(service, template)
	httpAdapter := fiber.NewFiberAdapter(service)

	fmt.Println("Server run in http://localhost:4321")

	httpAdapter.Run("localhost:4321")
}
