package domain

import "gorm.io/gorm"



type Admins struct{
	gorm.Model
    UserName   string `json:"username" gorm:"not null"`
	Password   string  `json:"password" gorm:"not null"`
}