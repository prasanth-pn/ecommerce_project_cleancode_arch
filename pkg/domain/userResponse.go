package domain





type UserResponse struct{
	ID           uint     `json:"id" gorm:"primaryKey"`
	First_Name   string 	`json:"first_name" gorm:"not null"`
	Last_Name    string 	`json:"last_name" gorm:"not null"`
	Email        string 	`json:"email" gorm:"not null" valid:"email"`
	Gender       string 	`json:"gender"`
	Phone        string 	`json:"phone" gorm:"not null"`
	Password     string 	`json:"password" gorm:"not null" valid:"length(6/12)"`
	Status       bool   	`json:"status"`
	Verification bool  		`json:"verification"`
	Token     	 string 		`json:"token"`

}
type AdminResponse struct {
	ID       int    `json:"id_login"`
	UserName string `json:"email"`
	Password string `json:"password,omitempty"`
	Token    string `json:"token,omitempty"`

}
