package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	"fmt"

	Mod "GoEcommerceProject/Models"
)

func AddBrand(brand Mod.Brand) bool{
	db := Cfg.DBConnect()
	db.Create(&brand)
	if brand.ID !=0 {
		return true
	}
	defer db.Close()
	return false
}

func Brands(brands []Mod.Brand, where ... interface{}) []Mod.Brand {
	db := Cfg.DBConnect()
	db.Find(&brands, where...)
	defer db.Close()
	return brands
}


func BrandWithProducts(brand Mod.Brand) Mod.Brand{
	db := Cfg.DBConnect()

	db.Find(&brand, ).Related(&brand.Products)


	defer db.Close()
	fmt.Println(brand)
	return brand
}


func UpdateBrand(brand Mod.Brand) bool {
	db := Cfg.DBConnect()
	if db.Save(&brand).RowsAffected == 1{
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}


func DeleteBrand(brand Mod.Brand) bool{
	db := Cfg.DBConnect()
	if db.Unscoped().Delete(&brand).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}
