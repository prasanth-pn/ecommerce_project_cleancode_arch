package repository

import (
	interfaces "clean/pkg/repository/interfaces"
	"database/sql"
)


type AdminDatabase struct{
	DB *sql.DB
}
func NewAdminRepository(DB *sql.DB)interfaces.AdminRepository{
	return &userDatabase{DB}
}

