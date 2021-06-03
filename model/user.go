package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name         string  `gorm:"type:varchar(20);not null"`       // string默认长度为255, 使用这种tag重设。
	Telephone          string     `gorm:"varchar(110;not null;unique"` // 自增
	Password 			string  `gorm:"varchar(110;not null "`
}