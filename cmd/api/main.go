package main

import (
	"log"

	_ "github.com/prasanth-pn/ecommerce_project_cleancode_arch/cmd/api/docs"
	_ "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/common/response"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/db"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/di"
)

// @title Go + Gin MobDex API
// @version 1.0
// @description PROJECT UPDATING YOU CAN VISIT GitHub repository AT https://github.com/prasanth-pn/ecommerce_project_cleancode_arch
// @description USED TECHNOLOGIES IN THE PROJECT IS VIPER, WIRE GIN-GONIC FRAMEWORK,  
// @description POSTGRESQL DATABASE, JWT AUTHENTICATION, DOCKER,DOCKER COMPOSE
// @contact.name PRASANTH P N
// @contact.url http://www.swagger.io/support
// @contact.email support@prasanthpn68@gmail.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {

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
