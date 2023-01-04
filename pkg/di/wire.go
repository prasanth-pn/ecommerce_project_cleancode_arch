//go:build wireinject
// +build wireinject

package di

import (
	http "clean/pkg/api"
	handler "clean/pkg/api/handler"
	config "clean/pkg/config"
	db "clean/pkg/db"
	usecase "clean/pkg/usecase"

	"github.com/google/wire"

	repository "clean/pkg/repository"
)

func InitializeEvent(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDB,
		 repository.NewUserRepository,
		  usecase.NewUserUseCase,
		  usecase.NewJWTService,
		 handler.NewUserHandler,
		  repository.NewAuthRepository,
		  usecase.NewAuthUseCase,
		 handler.NewAuthHandler,
		 http.NewServerHTTP,
		)
	return &http.ServerHTTP{}, nil
}
