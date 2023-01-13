package domain

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" validate:"required,min=3,max=50" gorm:"not null;unique"`
	Password string `json:"password" validate:"required,min=6,max=12" gorm:"not null" valid:"length(5/12)"`
}

type Users struct {
	User_Id      uint   `gorm:"primaryKey;autoIncrement:true;unique"`
	First_Name   string `json:"first_name" validate:"required,min=3,max=12" gorm:"not null"`
	Last_Name    string `json:"last_name" validate:"required,min=1,max=12" gorm:"not null"`
	Email        string `json:"email" validate:"required,min=3,max=50" gorm:"not null;unique"`
	Gender       string `json:"gender" validate:"required,min=4,max=8" `
	Phone        string `json:"phone" validate:"required,min=3,max=12" gorm:"not null;unique"`
	Password     string `json:"password" validate:"required,min=6,max=12" gorm:"not null" valid:"length(5/12)"`
	Status       bool
	Verification bool
	Country      string `json:"country "`
	City         string `json:"city " `
	Block_Status bool
	Cart_Id      uint
	Address_Id   uint
	Category_Id  uint
	//	Orders            Orders
	Orders_ID uint
	//WishList          WishList
	WishListID uint
	Created_At time.Time
	Updated_At time.Time
}
type Cart struct {
	Created_At  time.Time
	Deleted_At  time.Time
	Updated_At  time.Time
	Cart_Id     uint    ` gorm:"primaryKey;autoIncrement:true;unique"`
	User_Id     uint    `json:"user_id"   `
	Product_Id  uint    `json:"product_id" `
	Quantity    uint    `json:"quantity"`
	Total_Price float32 `json:"total_price"`
}
type Cartsinfo struct {
	gorm.Model
	User_Id      string `json:"user_id"`
	Product_Id   string `json:"product_id"`
	Product_Name string `json:"product_name"`
	Price        string `json:"price"`
	Email        string
	Quantity     string
	Total_Price  string
}

type Address struct {
	Address_id   uint   `json:"address_id" gorm:"primaryKey;unique"  `
	User_Id      uint   `json:"user_id"  gorm:"not null" `
	FName        string `json:"fname"  gorm:"not null" `
	LName        string `json:"lname" gorm:"not null"`
	Phone_Number int    `json:"phone_number"  gorm:"not null" `
	Pincode      int    `json:"pincode"  gorm:"not null" `
	House        string `json:"house" gorm:"not null"   `
	Area         string `json:"area" gorm:"not null"  `
	Landmark     string `json:"landmark"  gorm:"not null" `
	City         string `json:"city"  gorm:"not null" `
}
type PaymentMethod struct {
	COD bool
}
type Orders struct {
	gorm.Model
	User_Id         uint   `json:"user_id"  gorm:"not null" `
	Order_Id        string `json:"order_id"  gorm:"not null" `
	Total_Amount    uint   `json:"total_amount"  gorm:"not null" `
	Applied_Coupons string `json:"applied_coupons"  `
	Discount        uint   `json:"discount"   `
	PaymentMethod   string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status  string `json:"payment_status"   `
	Order_Status    string `json:"order_status"   `
	//Address         Address
	Address_Id uint `json:"address_id"  `
}
type Orderd_Items struct {
	gorm.Model
	User_Id         uint `json:"user_id"  gorm:"not null" `
	Product_id      uint `json:"product_id"`
	Order_Id        string
	Product_Name    string
	Price           string
	Order_Status    string
	Payment_Status  string
	PaymentMethod   string
	Applied_Coupons string
	Total_Amount    uint
}
type WishList struct {
	gorm.Model
	UserID     uint
	Product_Id uint
}
