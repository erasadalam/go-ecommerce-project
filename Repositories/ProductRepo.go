package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
	"github.com/jinzhu/gorm"
)

func AddProduct(product Mod.Product) bool{
	db := Cfg.DBConnect()
	db.Debug().Model(&product).Where("product_sl >= ?", product.ProductSL).Update("product_sl", gorm.Expr("product_sl + ?", 1))
	db.Debug().Create(&product)
	if product.ID !=0 {
		return true
	}
	defer db.Close()
	return false
}


func Product(product Mod.Product) Mod.Product {
	db := Cfg.DBConnect()
	db.Find(&product)
	defer db.Close()
	return product
}


func ProductsWithOthers(products []Mod.Product, where ... interface{}) []Mod.Product {
	db := Cfg.DBConnect()
	db.Order("product_sl ASC").Find(&products, where...)				//fetch all products(rows) from product table where status=1, and save into products address
	for i, _ := range products {
		db.Find(&products[i].Category, "id=?", products[i].CategoryID)		//fetch all from Category table of categoruID i,
		db.Find(&products[i].Brand, "id=?", products[i].BrandID)		//fetch all from Brand table of brandID i,
	}
	defer db.Close()
	return products
}


/*func UpdateProduct(product Mod.Product) bool {
	db := Cfg.DBConnect()
	var product_a Mod.Product
	var product_b Mod.Product
	db.Debug().Find(&product_a, "id = ?", product.ID)
	db.Debug().Find(&product_b, "product_sl = ?", product.ProductSL)
	db.Debug().Model(&product_b).Where("id = ?", product_b.ID).Update("product_sl", product_a.ProductSL)
	if db.Debug().Save(&product).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}*/

func UpdateProduct(product Mod.Product) bool {
	db := Cfg.DBConnect()
	var product_a Mod.Product
	var product_b Mod.Product
	var product_c Mod.Product
	db.Debug().Find(&product_a, "id = ?", product.ID)
	db.Debug().Find(&product_b, "product_sl = ?", product.ProductSL)
	//db.Debug().Model(&product_b).Where("id = ?", product_b.ID).Update("product_sl", product_a.ProductSL)

	if product.ProductSL < product_a.ProductSL {
		db.Debug().Model(&product_c).Where("product_sl >= ? and product_sl < ?", product.ProductSL, product_a.ProductSL).Update("product_sl", gorm.Expr("product_sl + ?", 1))
	}

	if product.ProductSL > product_a.ProductSL {
		db.Debug().Model(&product_c).Where("product_sl > ? and product_sl <= ?",  product_a.ProductSL, product.ProductSL).Update("product_sl", gorm.Expr("product_sl - ?", 1))
	}


	if db.Debug().Save(&product).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}


func DeleteProduct(product Mod.Product) bool{
	db := Cfg.DBConnect()
	if db.Unscoped().Delete(&product).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}
