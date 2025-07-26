package models

import (
	"errors"
	// "fmt"

	"github.com/AjayKumar-j-s/RecipeSharing/db"
	"github.com/AjayKumar-j-s/RecipeSharing/utils"
)

type User struct {
    UserID   int64  `gorm:"primaryKey;autoIncrement"`
	Name     string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}


func (u *User) Save()(error){


	Hpass,err := utils.HashPassword(u.Password)

	if(err != nil){
		return err
	}

	u.Password = Hpass


	res := db.DB.Create(u)


	return res.Error	
}

func (u *User)ValidatePAssword()(error){

	var u1 User
	err := db.DB.Select("user_id,password").Where("email = ?", u.Email).Find(&u1).Error

	if(err != nil){
		return err
	}
	// fmt.Print("result query",u1.UserID)
	Isvalid := utils.CheckHash(u.Password,u1.Password)

	u.UserID = u1.UserID

	if(!Isvalid){
		return errors.New("incorrect Password")
	}

	return nil



}