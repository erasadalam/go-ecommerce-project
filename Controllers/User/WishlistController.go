package User

import (
	G "GoEcommerceProject/Globals"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	S "GoEcommerceProject/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)



func AddtoWishList(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)

	if !authUser && !guest {
		return
	}


	if !authUser {
		G.Msg.Fail = "User must be logged in."
		session_wishlist, _ := G.Store.Get(c.Request, "login_token")
		session_wishlist.Values["req_url"] = c.Request.RequestURI
		session_wishlist.Save(c.Request, c.Writer)
		c.Redirect(http.StatusFound,"/login?q=" + session_wishlist.Values["req_url"].(string))
		return
	}

	if _, authAdminUser := M.IsAuthAdminUser(c, G.Store); authAdminUser {
		return
	}


	var wishlist Mod.Wishlist
	id, _ := strconv.Atoi(c.Param("id"))
	wishlist.ProductID = uint(id)
	wishlist.UserID = user.ID
	_, success := R.WishlistExistProduct(wishlist)
	if success {
		G.Msg.Fail = "This product already added into your wishlist."

		c.Redirect(http.StatusFound, "/")
		return
	}
	if R.Wishlist(wishlist) {
		G.Msg.Success = "Product Added into wishlist."
		c.Redirect(http.StatusFound, "/")
		return
	} else {
		G.Msg.Fail = "Some Error Occured. Please Try Again Later."
		c.Redirect(http.StatusFound, "/")
		return
	}

	/*c.HTML(http.StatusOK, "cat-wise-pro.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Category-Wise-Product", "Category": nil})*/
}

func AddCartAndRemove(c * gin.Context) {
	_, authUser := M.IsGuest(c, G.Store)
	_, guest := M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	success := S.AddCart(c, id, quantity)
	if success {
		DeleteFromWishList(c, id)
		G.Msg.Success = "Product added into cart and deleted from your Wishlist."
		c.Redirect(http.StatusFound, "/wish-list")

	} else {
		G.Msg.Fail = "Some Error Occurred, Please Try Again."
		c.Redirect(http.StatusFound, "/")
	}
}

func DeleteFromWishList(c *gin.Context, id int) {
	var user Mod.User
	var authUser bool
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser  {
		G.Msg.Fail = "User must loged in."
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var wishlist Mod.Wishlist
	//id, _ := strconv.Atoi(c.Param("id"))
	wishlist.ProductID = uint(id)
	wishlist.UserID = user.ID
	success := R.DeleteProductFromWishlist(wishlist, wishlist.ProductID, wishlist.UserID)
	if success {
		return
	}

}

func SingleDeleteFromWishList(c *gin.Context) {
	var user Mod.User
	var authUser bool
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser  {
		G.Msg.Fail = "User must loged in."
		c.Redirect(http.StatusFound, "/login")
		return
	}
	var wishlist Mod.Wishlist
	id,_ := strconv.Atoi(c.Param("id"))
	wishlist.ProductID = uint(id)
	wishlist.UserID = user.ID
	success := R.DeleteProductFromWishlist(wishlist, wishlist.ProductID, wishlist.UserID)
	if success {
		G.Msg.Success = "Product deleted from wishlist"
		c.Redirect(http.StatusFound, "/wish-list")
	}
}

func ShowWishList(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}


	 if !authUser {
		 G.Msg.Fail = "User must be logged in."
		 session_wishlist, _ := G.Store.Get(c.Request, "login_token")
		 session_wishlist.Values["req_url"] = c.Request.RequestURI
		 session_wishlist.Save(c.Request, c.Writer)
		 c.Redirect(http.StatusFound,"/login?q=" + session_wishlist.Values["req_url"].(string))
		 return
	}


	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)
	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]


	session, _ := G.Store.Get(c.Request, "cart")
	cart_items := len(session.Values)

	c.HTML(http.StatusOK, "wish-list.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Wishlist",  "Wishlists": wishlists, "Cart_items": cart_items, "Wish_items": session_wishlist.Values["wish_items"]})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}