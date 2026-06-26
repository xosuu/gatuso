package model

import (
	"gorm.io/gorm"
)



type User struct{
	gorm.Model
	Name string
	Email string
	Movie []Movie
	
}


func(u *User) GetId()uint{
	return u.ID
}
