package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
)

func ConnectDB(cfg config.Config) *sql.DB {
	databaseName := cfg.DBName
	dbURI := cfg.DBSOURCE
	db, err := sql.Open("postgres", dbURI)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error in ping")
		log.Fatal("this errror is :",err)
	}
	log.Println("\n Connected to database:", databaseName)
	return db
}
