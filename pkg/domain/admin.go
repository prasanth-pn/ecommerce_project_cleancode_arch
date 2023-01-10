package domain

import "gorm.io/gorm"

type Admins struct {
	gorm.Model
	UserName string `json:"username" gorm:"not null"`
	Password string `json:"password" gorm:"not null"`
}
type Category struct {
	Category_id   int    `gorm:"primaryKey;"`
	Category_Name string `json:"category_name" gorm:"not null"`
	Description   string `json:"description" gorm:"not null"`
	Image         string `json:"image"`
}

type Product struct {
	Product_Id   int     `gorm:"primaryKey"`
	Product_name string  `json:"product_name" gorm:"not null"`
	Description  string  `json:"description" gorm:"not null"`
	Quantity     uint16  `json:"quantity" gorm:"not null"`
	Image_Path   string  `json:"image_path" gorm:"not null"`
	Price        float32 `json:"price" gorm:"not null"`
	Color        string  `json:"color"`
	Available    bool    `json:"available" gorm:"not null"`
	Trending     bool    `json:"trending" gorm:"not null"`
	Category_id  uint    `json:"category_id"`
	Cart        Cart
	Cart_id     uint `json:"cart_id" `
	//Category    Category
	//WishList    WishList
	WishListID  uint
}
