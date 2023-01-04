package db

import (
	"clean/pkg/config"
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg config.Config)(*sql.DB){
	databaseName:=cfg.DBName
	dbURI:=cfg.DBSOURCE
    db,err:=sql.Open("postgres",dbURI)

	if err!=nil{
		log.Fatal(err)
	}

	err=db.Ping()
	if err!=nil{
		fmt.Println("error in ping")
		log.Fatal(err)
	}
	log.Println("\n Connected to database:",databaseName)
	return db
}