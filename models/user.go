package models

import "github.com/AjayKumar-j-s/RecipeSharing/db"

type User struct {
	UserID   int64  `gorm:"primaryKey"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}


func (u *User) Save()(error){
	res := db.DB.Create(u)


	return res.Error	
}
