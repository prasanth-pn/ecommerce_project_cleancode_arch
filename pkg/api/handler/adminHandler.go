package handler

import (
	services "clean/pkg/usecase/interfaces"
)

type AdminHandler struct {
	adminService services.AdminUseCase
	jwtService   services.JWTService
}

func NewAdminHandler(usecase services.AdminUseCase, jwtService services.JWTService) *AdminHandler {
	return &AdminHandler{
		adminService: usecase,
		jwtService:   jwtService,
	}
}
