package domain

import (
	"time"

	"gorm.io/gorm"
)

type Login struct {
	Email    string `json:"email" validate:"required,min=3,max=50" gorm:"not null;unique"`
	Password string `json:"password" validate:"required,min=6,max=12" gorm:"not null" valid:"length(5/12)"`
}
type Verification struct {
	Creat_At time.Time
	Exp_At   time.Time
	Email    string `json:"email" gorm:"not null;unique"`
	Code     int    `json:"code" gorm:"not null"`
}
type Users struct {
	User_Id      uint   `gorm:"serial primaryKey;autoIncrement:true;unique"`
	First_Name   string `json:"first_name" validate:"required,min=3,max=12" gorm:"not null"`
	Last_Name    string `json:"last_name" validate:"required,min=1,max=12" gorm:"not null"`
	Email        string `json:"email" validate:"required,min=3,max=50" gorm:"not null;unique"`
	Gender       string `json:"gender" validate:"required,min=4,max=8" `
	Phone        string `json:"phone" validate:"required,min=3,max=12" gorm:"not null;unique"`
	Password     string `json:"password" validate:"required,min=6,max=12" gorm:"not null" valid:"length(5/12)"`
	Verification bool
	Country      string `json:"country "`
	City         string `json:"city " `
	Block_Status bool
	Created_At   time.Time
	Updated_At   time.Time
	Profile_Pic  string
}
type Cart struct {
	Created_At  time.Time
	Updated_At  time.Time
	Cart_Id     uint ` gorm:" serial primaryKey;autoIncrement:true;unique"`
	Image       string
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
	Address_id   uint   ` gorm:" serial primaryKey;autoIncrement:true;unique"`
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
	Created_at      time.Time
	Canceled_at     time.Time
	User_Id         uint   `json:"user_id"  gorm:"not null" `
	Order_Id        string `json:"order_id"  gorm:"not null" `
	Applied_Coupons string `json:"applied_coupons"  `
	Discount        uint   `json:"discount"`
	Total_Amount    uint   `json:"total_amount"  gorm:"not null" `
	Balance_Amount  int    `json:"balance_amount"`
	PaymentMethod   string `json:"paymentmethod"  gorm:"not null" `
	Payment_Status  string `json:"payment_status"   `
	Payment_Id      string `json:"payment_id"`
	Order_Status    string `json:"order_status"`
	Address_Id      uint   `json:"address_id" `
}
type Ordered_Items struct {
	Product_id   uint   `json:"product_id"`
	Order_Id     string `json:"order_id"`
	Product_Name string
	Image_Path   string
	Price        int
	Quantity     uint
}

type WishList struct {
	gorm.Model
	UserID     uint
	Product_Id uint
}
type ListOrder struct {
	Created_At      time.Time
	Order_id        string
	Total_amount    int
	Applied_Coupons string
	Discount        uint
	Payment_Method  string
	Payment_Status  string
	Payment_id      string
	Order_Status    string
	Address_Id      uint
	Product_Id      uint
	Product_Name    string
	Image_Path      string
	Price           int
	Quantity        uint
}
type Applied_Coupons struct {
	Created_At  time.Time
	UserID      uint
	Coupon_Code string `json:"coupon_code"`
}
