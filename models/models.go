package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserId  int     `json:"UserId" gorm:"int;not null"`
	Balance float64 `json:"balance" gorm:"float;default:0"`
}
