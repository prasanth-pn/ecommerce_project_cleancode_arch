package domain

import "github.com/golang-jwt/jwt/v4"





type UserResponse struct{
	ID           uint     `json:"id"`
	First_Name   string 	`json:"first_name"`
	Last_Name    string 	`json:"last_name" `
	Email        string 	`json:"email" gorm:"not null"`
	Gender       string 	`json:"gender"`
	Phone        string 	`json:"phone" gorm:"not null"`
	Password     string 	`json:"password" gorm:"not null" valid:"length(6/12)"`
	Status       bool   	`json:"status"`
	Verification bool  		`json:"verification"`
	Token     	 string 	`json:"token"`

}
type AdminResponse struct {
	ID       int    `json:"id_login"`
	UserName string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`

}
type SignedDetails struct {
	User_Id    uint
	First_Name string
	Email      string
	jwt.StandardClaims
}


type ProductResponse struct{
	Product_Id    int
	Product_Name  string
	Description   string
	Image         string
	Price         float32
	Color         string
	Available     bool
	Trending      bool
	Category_Name string
}
