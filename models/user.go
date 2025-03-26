package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"-"` // ห้ามส่งค่า password ออกไป
	Name     string `json:"name"`
	Address  string `json:"address"`
}
