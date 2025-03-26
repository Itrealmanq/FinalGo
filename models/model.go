package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductID     uint       `gorm:"primaryKey;column:product_id" json:"product_id"`
	ProductName   string     `gorm:"column:product_name;size:255" json:"product_name"`
	Description   string     `gorm:"column:description" json:"description"`
	Price         float64    `gorm:"column:price;type:decimal(10,2)" json:"price"`
	StockQuantity int        `gorm:"column:stock_quantity" json:"stock_quantity"`
	CartItems     []CartItem `gorm:"foreignKey:ProductID" json:"-"`
}

type Customer struct {
	gorm.Model
	CustomerID  uint   `gorm:"primaryKey;column:customer_id" json:"customer_id"`
	FirstName   string `gorm:"column:first_name;size:255" json:"first_name"`
	LastName    string `gorm:"column:last_name;size:255" json:"last_name"`
	Email       string `gorm:"column:email;size:255" json:"email"`
	PhoneNumber string `gorm:"column:phone_number;size:20" json:"phone_number"`
	Address     string `gorm:"column:address;size:255" json:"address"`
	Password    string `gorm:"column:password;size:255" json:"password,omitempty"`
	Carts       []Cart `gorm:"foreignKey:CustomerID" json:"-"`
}

type Cart struct {
	gorm.Model
	CartID     uint       `gorm:"primaryKey;column:cart_id" json:"cart_id"`
	CustomerID uint       `gorm:"column:customer_id" json:"customer_id"`
	CartName   string     `gorm:"column:cart_name;size:255" json:"cart_name"`
	Customer   Customer   `gorm:"foreignKey:CustomerID" json:"-"`
	CartItems  []CartItem `gorm:"foreignKey:CartID" json:"cart_items"`
}

type CartItem struct {
	gorm.Model
	CartItemID uint    `gorm:"primaryKey;column:cart_item_id" json:"cart_item_id"`
	CartID     uint    `gorm:"column:cart_id" json:"cart_id"`
	ProductID  uint    `gorm:"column:product_id" json:"product_id"`
	Quantity   int     `gorm:"column:quantity" json:"quantity"`
	Cart       Cart    `gorm:"foreignKey:CartID" json:"-"`
	Product    Product `gorm:"foreignKey:ProductID" json:"-"`
}

// TableName sets the table name for each model
func (Product) TableName() string {
	return "product"
}

func (Customer) TableName() string {
	return "customer"
}

func (Cart) TableName() string {
	return "cart"
}

func (CartItem) TableName() string {
	return "cart_item"
}
