//go:build wireinject
// +build wireinject

package di

import (
	http "clean/pkg/api"
	handler "clean/pkg/api/handler"
	config "clean/pkg/config"
	db "clean/pkg/db"
	usecase "clean/pkg/usecase"
	middleware"clean/pkg/api/middleware"

	"github.com/google/wire"

	repository "clean/pkg/repository"
)

func InitializeEvent(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDB,
		 repository.NewUserRepository,
		 repository.NewAdminRepository,
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

		)
	return &http.ServerHTTP{}, nil
}
