// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api/handler"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api/middleware"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/db"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/repository"
	"github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/usecase"
)

// Injectors from wire.go:

func InitializeEvent(cfg config.Config) (*http.ServerHTTP, error) {
	sqlDB := db.ConnectDB(cfg)
	userRepository := repository.NewUserRepository(sqlDB)
	userUseCase := usecase.NewUserUseCase(userRepository)
	authRepository := repository.NewAuthRepository(sqlDB)
	mailConfig := config.NewMailConfig()
	authUseCase := usecase.NewAuthUseCase(authRepository, mailConfig, cfg)
	userHandler := handler.NewUserHandler(userUseCase, authUseCase)
	jwtService := usecase.NewJWTService()
	authHandler := handler.NewAuthHandler(authUseCase, jwtService)
	adminRepository := repository.NewAdminRepository(sqlDB)
	adminUseCase := usecase.NewAdminUseCase(adminRepository)
	adminHandler := handler.NewAdminHandler(adminUseCase, jwtService, authUseCase)
	middlewareMiddleware := middleware.NewUserMiddleware(jwtService)
	serverHTTP := http.NewServerHTTP(userHandler, authHandler, adminHandler, middlewareMiddleware)
	return serverHTTP, nil
}
