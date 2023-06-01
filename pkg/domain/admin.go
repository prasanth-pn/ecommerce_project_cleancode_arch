package domain

import "time"

type Admins struct {
	ID       uint   `gorm:"primaryKey;autoIncrement:true;unique"`
	UserName string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
type Coupon struct {
	Created_At time.Time
	Coupon_Id  uint   `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Coupon     string `json:"coupon"`
	Discount   int    `json:"discount"`
	Quantity   int    `json:"quantity"`
	Validity   int64  `json:"validity"`
}
type Category struct {
	Category_Id   int    `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Category_Name string `json:"category_name" gorm:"not null;unique"`
	Description   string `json:"description" gorm:"not null"`
	Image         string `json:"_"`
}

type Product struct {
	Product_Id   int    `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Product_Name string `json:"product_name" gorm:"not null;unique"`
	Image        string
	Description  string  `json:"description" gorm:"not null"`
	Quantity     uint16  `json:"quantity" gorm:"not null"`
	Price        float32 `json:"price" gorm:"not null"`
	Color        string  `json:"color"`
	Available    bool    `json:"available" gorm:"not null"`
	Trending     bool    `json:"trending" gorm:"not null"`
	Category_Id  uint    `json:"category_id"`
	Brand_Id     uint    `json:"brand_id"`
}
type Brand struct {
	Brand_Id          uint   `gorm:"serial primaryKey;autoIncrement:true;unique"`
	Brand_Name        string `json:"brand_name" gorm:"not null;unique"`
	Brand_Description string `json:"brand_description" gorm:"not null"`
	Discount          uint   `json:"discount"`
}

type Model struct {
	Model_Id   uint   `json:"model_id" gorm:"primaryKey;autoIncrement:true;unique"`
	Model_Name string `json:"model_name" gorm:"not null;unique"`
}
