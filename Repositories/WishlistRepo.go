package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
	_ "github.com/go-sql-driver/mysql"
)






func WishlistExistProduct(wishlist Mod.Wishlist) (Mod.Wishlist, bool) {
	db := Cfg.DBConnect()
	notFound := db.First(&wishlist, "user_id=? and product_id=?", wishlist.UserID,wishlist.ProductID).RecordNotFound()

	if notFound {
		defer db.Close()
		return wishlist, false
	} else {
		defer db.Close()
		return wishlist, true
	}
}


func DeleteProductFromWishlist(wishlist Mod.Wishlist, ProductID uint, UserID uint) bool {
	db := Cfg.DBConnect()
	if db.Debug().Unscoped().Delete(&wishlist, "user_id=? and product_id=?",UserID, ProductID).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}

func Wishlist(wishlist Mod.Wishlist) bool{
	db := Cfg.DBConnect()
	db.Debug().Create(&wishlist)
	if wishlist.UserID != 0 {
		defer db.Close()
		return true
	}
	defer db.Close()
	return false
}



func WishlistWithProducts(wishlists []Mod.Wishlist, userID uint) []Mod.Wishlist {
	db := Cfg.DBConnect()
	db.Debug().Find(&wishlists, "user_id=?", userID)

	for i,_ := range wishlists {
		db.Debug().Find(&wishlists[i]).Related(&wishlists[i].Product)

	}


	defer db.Close()
	return wishlists
}
