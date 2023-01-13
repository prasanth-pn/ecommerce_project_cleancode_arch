package usecase

import "clean/pkg/domain"

func (c * userUseCase)AddAddress(address domain.Address)error{
	err:=c.userRepo.AddAddress(address)
	return err
}
func( c *userUseCase)ListAddress(user_id uint)([]domain.Address,error){
	address,err:=c.userRepo.ListAddress(user_id)
	return address ,err
}

