package domain

import "gorm.io/gorm"


type Users struct{
	gorm.Model
	First_Name   string `json:"first_name" gorm:"not null"`
	Last_Name    string `json:"last_name" gorm:"not null"`
	Email        string `json:"email" gorm:"not null"`
	Gender       string `json:"gender"`
	Phone        string `json:"phone" gorm:"not null"`
	Password     string `json:"password" gorm:"not null" valid:"length(5/12)"`
	Status       bool   `json:"status"`
	Verification bool   `json:"verification"`
}