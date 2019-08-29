package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//Product struct
type Product struct {
	gorm.Model
	Image       string  `json:"image"`
	ImgAlt      string  `gorm:"column:imgalt" json:"imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"` //sql.NullFloat64
	ProductName string  `gorm:"column:productname" json:"productname"`
	Description string  `json:"description"`
}

//TableName declares table name
func (Product) TableName() string {
	return "products"
}

//Customer struct
type Customer struct {
	gorm.Model
	FirstName string  `gorm:"column:firstname" json:"firstname"`
	LastName  string  `gorm:"column:lastname" json:"lastname"`
	Email     string  `gorm:"column:email" json:"email"`
	Pass      string  `json:"password"`
	LoggedIn  bool    `gorm:"column:loggedin" json:"loggedin"`
	Orders    []Order `gorm:"foreignkey:CustomerID" json:"orders"`
}

//TableName declares table name
func (Customer) TableName() string {
	return "customers"
}

//Order struct
type Order struct {
	gorm.Model
	CustomerID   int       `gorm:"column:customer_id"`
	ProductID    int       `gorm:"column:product_id"`
	Price        float64   `gorm:"column:sell_price"`
	PurchaseDate time.Time `gorm:"column:purchase_date" json:"purchase_date"`
}

//TableName declares table name
func (Order) TableName() string {
	return "orders"
}
