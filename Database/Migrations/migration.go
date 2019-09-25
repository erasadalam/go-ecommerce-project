package Migrtaions

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
	"github.com/jinzhu/gorm"
)


func Migrate() {
	db := Cfg.DBConnect()
	db.AutoMigrate(&Mod.Role{})
	db.AutoMigrate(&Mod.User{})
	db.AutoMigrate(&Mod.Category{})
	db.AutoMigrate(&Mod.Brand{})
	db.AutoMigrate(&Mod.Product{})
	db.AutoMigrate(&Mod.Bill{})
	db.AutoMigrate(&Mod.PayMethod{})
	db.AutoMigrate(&Mod.Order{})
	db.AutoMigrate(&Mod.OrderDetail{})
	db.AutoMigrate(&Mod.Wishlist{})
	AddForeignKeys(db)
	defer db.Close()
}

func AddForeignKeys(db *gorm.DB) {
	db.Model(&Mod.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.Product{}).AddForeignKey("category_id", "categories(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.Product{}).AddForeignKey("brand_id", "brands(id)", "RESTRICT", "RESTRICT")

	db.Model(&Mod.Bill{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.Order{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.Order{}).AddForeignKey("bill_id", "bills(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.Order{}).AddForeignKey("pay_method_id", "pay_methods(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.OrderDetail{}).AddForeignKey("order_id", "orders(id)", "RESTRICT", "RESTRICT")
	db.Model(&Mod.OrderDetail{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

	//db.Model(&Mod.Wishlist{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	//db.Model(&Mod.Wishlist{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")

}


