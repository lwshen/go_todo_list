package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username       string `gorm:"unique"`
	PasswordDigest string //存储的是加密后的密码
}
