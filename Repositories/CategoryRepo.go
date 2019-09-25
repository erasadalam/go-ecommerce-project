package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
)

func AddCategory(category Mod.Category) bool {
	db := Cfg.DBConnect()
	db.Create(&category)
	if category.ID !=0 {
		return true
	}
	defer db.Close()
	return false
}


func  Categories(categories []Mod.Category, where ... interface{}) []Mod.Category {
	db := Cfg.DBConnect()
	db.Find(&categories, where...)
	defer db.Close()
	return categories
}


func CategoryWithProducts(category Mod.Category) Mod.Category{
	db := Cfg.DBConnect()
	db.Find(&category).Related(&category.Products)
	defer db.Close()
	return category
}


func UpdateCategory(category Mod.Category) bool {
	db := Cfg.DBConnect()
	if db.Save(&category).RowsAffected == 1{
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}


func DeleteCategory(category Mod.Category) bool{
	db := Cfg.DBConnect()
	if db.Unscoped().Delete(&category).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}
