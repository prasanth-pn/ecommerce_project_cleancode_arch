//go:build wireinject
// +build wireinject

package di

import (
	http "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api"
	handler "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api/handler"
	middleware "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/api/middleware"
	config "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/config"
	db "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/db"
	usecase "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/usecase"

	"github.com/google/wire"

	repository "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/repository"
)

func InitializeEvent(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(
		repository.NewUserRepository,
		repository.NewAdminRepository,
		config.NewMailConfig,
		usecase.NewUserUseCase,
		usecase.NewAdminUseCase,
		usecase.NewAuthUseCase,
		usecase.NewJWTService,
		handler.NewUserHandler,
		handler.NewAdminHandler,
		repository.NewAuthRepository,
		handler.NewAuthHandler,
		middleware.NewUserMiddleware,
		http.NewServerHTTP,
		db.ConnectDB,
	)
	return &http.ServerHTTP{}, nil
}
