package entities

import (
	"gorm.io/gorm"
)
type User struct{
	gorm.Model
	Id uint
	Username string
	Password string
}
func (User) TableName() string{
	return "User"
}