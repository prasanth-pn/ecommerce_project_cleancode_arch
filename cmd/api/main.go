package main

import (
	"log"

	_ "github.com/prasanth-pn/ecommerce_project_cleancode_arch/cmd/api/docs"
	_ "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/db"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/di"

	"github.com/joho/godotenv"
)

// @title Go + Gin ecommerce API
// @version 1.0
// @description This is a sample server  server. You can visit the GitHub repository at https://github.com/prasanth-pn/github.com/prasanth-pn/ecommerce_project_cleancode_arch-code-architecture-ecommerce

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading env file")
	}

	config, configErr := config.LoadConfig()
	if configErr != nil {
		log.Fatal("cannot load config :", configErr)

	}
	//connect database
	_, _ = db.ConnectGormDB(config)
	//fmt.Println(gorm)
	server, diErr := di.InitializeEvent(config)

	if diErr != nil {

		log.Fatal("cannot start server :", diErr)
	} else {
		server.Start()

	}

}
