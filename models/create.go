package models

import "github.com/AjayKumar-j-s/RecipeSharing/db"

func CreateTable() {

	db.DB.AutoMigrate(&User{})

	db.DB.AutoMigrate(&Recipe{})
}