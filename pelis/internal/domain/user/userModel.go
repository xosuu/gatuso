package user

import (
	"gorm.io/gorm"
)



type User struct{
	gorm.Model
	Name string
	Email string
	
}


func(u *User) GetId()uint{
	return u.ID
}
