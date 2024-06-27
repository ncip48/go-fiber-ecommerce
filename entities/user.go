// @/entities/dog.go
package entities

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
	IsActive bool   `json:"is_active" gorm:"default:true"`
}

type Profile struct {
	ID       uint   `gorm:"primarykey"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   int    `json:"gender"`
	IsActive bool   `json:"is_active"`
}
