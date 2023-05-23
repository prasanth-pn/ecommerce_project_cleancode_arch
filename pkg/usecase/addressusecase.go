package usecase

import "github.com/prasanth-pn/ecommerce_project_cleancode_arch/pkg/domain"

func (c *userUseCase) AddAddress(address domain.Address) error {
	err := c.userRepo.AddAddress(address)
	return err
}
func (c *userUseCase) ListAddress(user_id uint) ([]domain.Address, error) {
	address, err := c.userRepo.ListAddress(user_id)
	return address, err
}
