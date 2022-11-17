package models

import "gorm.io/gorm"

type UserBalance struct {
	gorm.Model
	UserId  int     `json:"UserId" gorm:"int;not null;unique"`
	Balance float64 `json:"balance" gorm:"float;default:0"`
}

type UserOrder struct {
	gorm.Model
	UserId     int     `json:"UserId" gorm:"int;not null"`
	ServiceId  int     `json:"ServiceId" gorm:"int;not null"`
	OrderId    int     `json:"OrderId" gorm:"int;not null"`
	Cost       float64 `json:"Cost" gorm:"float;default:0"`
	IsReserved bool    `json:"IsReserved" gorm:"bool;default:true"`
}
