package repository

import (
	//domain "clean/pkg/domain"
	interfaces "clean/pkg/repository/interfaces"
	
	"database/sql"

	//"gorm.io/gorm"
)
type userDatabase struct{
	DB *sql.DB
}
func NewUserRepository(DB *sql.DB)interfaces.UserRepository{
	return &userDatabase{DB}
}




// func (c *userDatabase)FindAll(ctx context.Context)([]domain.Users,error)  {

// 	return nil,nil
// }
	


