package models

import (
	"github.com/AjayKumar-j-s/RecipeSharing/db"
)

type Recipe struct {
    ID   		int64  `gorm:"primaryKey;autoIncrement"`
	Title       string
	Description string
	Instruction string
	UserID      int64
	User        User `gorm:"foreignKey:UserID"`
}

func (r *Recipe) Save()(error) {

	res := db.DB.Create(r)

	return res.Error

}

func GetAllrecipes()([]Recipe,error){

	var r []Recipe

	res := db.DB.Find(&r)

	if res.Error!= nil{
		return nil,res.Error
	}
	return r,nil
}


func GetById(id int64)(Recipe,error){

	var r Recipe

	res := db.DB.First(&r,id)

	if(res.Error != nil){
		return Recipe{},res.Error
	}

	return r,nil


}

func (r *Recipe)UpdateRec()(error){

	res := db.DB.Model(&Recipe{}).Where("id = ?", r.ID).Updates(map[string]interface{}{
  	"title": r.Title,
 	"description": r.Description,
  	"instruction": r.Instruction,
})

	if(res.Error != nil){
		return res.Error
	}
	return nil
}


func (r *Recipe)DeleteRec()(error){


	res := db.DB.Delete(&Recipe{},r.ID)

	if(res.Error!=nil){
		return res.Error
	}
	return nil
}	
