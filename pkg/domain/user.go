package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	User_Id      uint   `gorm:"primaryKey;autoIncrement:true;unique"`
	First_Name   string `json:"first_name" validate:"required,min=3,max=12" gorm:"not null"`
	Last_Name    string `json:"last_name" validate:"required,min=3,max=12" gorm:"not null"`
	Email        string `json:"email" validate:"required,min=3,max=12" gorm:"not null;unique"`
	Gender       string `json:"gender" validate:"required,min=3,max=12" `
	Phone        string `json:"phone" validate:"required,min=3,max=12" gorm:"not null"`
	Password     string `json:"password" validate:"required,min=3,max=12" gorm:"not null" valid:"length(5/12)"`
	Status       bool   `json:"status"`
	Verification bool   `json:"verification"`
	Country      string `json:"country "`
	City         string `json:"city "   `
	Pincode      uint   `json:"pincode " `
	Cart         Cart
	Cart_id      uint `json:"cart_id" `
	Address        Address
	Address_id uint `json:"address_id" `
	Orders         Orders
	Orders_ID uint `json:"orders_id" `
	//Wallet_Balance uint `json:"wallet_balance" `
	//Applied_Coupons   Applied_Coupons
	Applied_CouponsID uint
	WishList          WishList
	WishListID uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
type Cart struct {
	Cart_id     uint `json:"cart_id" gorm:"primaryKey;autoIncrement:true;unique"  `
	User_Id     uint `json:"user_id"   `
	ProductID   uint `json:"product_id"  `
	Quantity    uint `json:"quantity" `
	Total_Price float32 `json:"total_price"   `
}
type Cartsinfo struct {
	gorm.Model
	User_id      string
	Product_id   string
	Product_Name string
	Price        string
	Email        string
	Quantity     string
	Total_Price  string
}
type Address struct {
	Address_id   uint   `json:"address_id" gorm:"primaryKey"  `
	UserId       uint   `json:"user_id"  gorm:"not null" `
	Name         string `json:"name"  gorm:"not null" `
	Phone_number int    `json:"phone_number"  gorm:"not null" `
	Pincode      int    `json:"pincode"  gorm:"not null" `
	House        string `json:"house"   `
	Area         string `json:"area"   `
	Landmark     string `json:"landmark"  gorm:"not null" `
	City         string `json:"city"  gorm:"not null" `
}
type PaymentMethod struct {
	COD bool
}
type Orders struct {
	gorm.Model
	UserId          uint   `json:"user_id"  gorm:"not null" `
	Order_id        string `json:"order_id"  gorm:"not null" `
	Total_Amount    uint   `json:"total_amount"  gorm:"not null" `
	Applied_Coupons string `json:"applied_coupons"  `
	Discount        uint   `json:"discount"   `
	PaymentMethod   string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status  string `json:"payment_status"   `
	Order_Status    string `json:"order_status"   `
	Address         Address
	Address_id      uint `json:"address_id"  `
}
type Orderd_Items struct {
	gorm.Model
	UserId          uint `json:"user_id"  gorm:"not null" `
	Product_id      uint
	OrdersID        string
	Product_Name    string
	Price           string
	Order_Status    string
	Payment_Status  string
	PaymentMethod   string
	Applied_Coupons string
	Total_amount    uint
}
type WishList struct {
	gorm.Model
	UserID     uint
	Product_id uint
}
