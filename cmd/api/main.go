package main

import (
	_ "clean/cmd/api/docs"
	"clean/pkg/config"
	"clean/pkg/db"
	"clean/pkg/di"
	"fmt"
	"log"
	"github.com/joho/godotenv"
)

// @title Go + Gin ecommerce API
// @version 1.0
// @description This is a sample server  server. You can visit the GitHub repository at https://github.com/LordGhostX/swag-gin-demo

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	err := godotenv.Load()
	if err!=nil{
		log.Fatal("error loading env file")
	}


	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config :", configErr)

	}
	//connect database
	gorm,_:=db.ConnectGormDB(config)
	fmt.Printf("\n\n%v",gorm)

	server, diErr := di.InitializeEvent(config)

	if diErr != nil {

		log.Fatal("cannot start server :", diErr)
	} else {
		server.Start()

	}

}
