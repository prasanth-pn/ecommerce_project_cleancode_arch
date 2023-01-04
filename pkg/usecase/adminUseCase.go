package usecase

import (
	
	interfaces "clean/pkg/repository/interfaces"
	services "clean/pkg/usecase/interfaces"
)



type adminUseCase struct{
	adminRepo interfaces.AdminRepository

}
func NewAdminUseCase(adminRepo interfaces.AdminRepository)services.AdminUseCase{
	return&adminUseCase{
		adminRepo: adminRepo,

	}
}