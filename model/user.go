package model

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	ID string `gorm:primarykey`
	Name string
	Email string
}