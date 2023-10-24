package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
	Avatar   string
}

type RegisterSturct struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Fullname string `json:"fullname"  binding:"required"`
	Avatar   string `json:"avatar"    binding:"required"`
}

type LoginStruct struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
