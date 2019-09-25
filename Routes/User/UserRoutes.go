package User

import (
	"GoEcommerceProject/Controllers/User"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.GET("/", User.Welcome)
	r.GET("/product-details/:id", User.ProductDetails)
	r.GET("/brand-wise-pro/:id", User.BrandWisePro)
	r.GET("/cat-wise-pro/:id", User.CatWisePro)
	r.GET("/add-to-wishlist/:id", User.AddtoWishList)
	r.POST("/add-cart", User.AddCart)
	r.POST("/add-cart-redirect-cartpage", User.AddCartRedirectCartpage)
	r.POST("/add-cart-and-remove", User.AddCartAndRemove)
	r.GET("/show-cart", User.ShowCart)
	r.GET("/wish-list", User.ShowWishList)
	r.GET("/delete-from-cart/:id", User.DeleteFromCart)
	r.GET("/delete-from-wishlist/:id", User.SingleDeleteFromWishList)
	r.GET("/checkout/:total", User.Checkout)
	r.GET("/bill-to", User.BillToGet)
	r.POST("/bill-to", User.BillToPost)

}
