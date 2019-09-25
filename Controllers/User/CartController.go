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


func AddCart(c * gin.Context) {
	_, authUser := M.IsGuest(c, G.Store)
	_, guest := M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	success := S.AddCart(c, id, quantity)
	if success {

		c.Redirect(http.StatusFound, "/")
	} else {
		G.Msg.Fail = "Some Error Occurred, Please Try Again."
		c.Redirect(http.StatusFound, "/")
	}
}



func AddCartRedirectCartpage(c * gin.Context) {
	_, authUser := M.IsGuest(c, G.Store)
	_, guest := M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	quantity, _ := strconv.Atoi(c.PostForm("quantity"))
	success := S.AddCart(c, id, quantity)
	if success {
		c.Redirect(http.StatusFound, "/show-cart")
	} else {
		G.Msg.Fail = "Some Error Occurred, Please Try Again."
		c.Redirect(http.StatusFound, "/")
	}
}


func ShowCart(c *gin.Context) {
	var user Mod.User
	//var authUser, guest bool
	//user, _ = M.IsGuest(c, G.Store)
	user, _ = M.IsAuthUser(c, G.Store)

/*	if !authUser {
		G.Msg.Fail = "User must be logged in."
		session_wishlist, _ := G.Store.Get(c.Request, "login_token")
		session_wishlist.Values["req_url"] = c.Request.RequestURI
		session_wishlist.Save(c.Request, c.Writer)
		c.Redirect(http.StatusFound,"/login?q=" + session_wishlist.Values["req_url"].(string))
		return
	}*/

/*	if !authUser && !guest {
		return
	}*/

	finalCart := S.ProcessCart(c)

	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)

	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]

	session, _ := G.Store.Get(c.Request, "cart")
	cart_items := len(session.Values)
	c.HTML(http.StatusOK, "show-cart.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Cart", "FinalCart":finalCart, "Cart_items": cart_items, "Wish_items": session_wishlist.Values["wish_items"]})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}




func DeleteFromCart(c *gin.Context) {
/*	_, authUser := M.IsGuest(c, G.Store)
	_, guest := M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}*/


	//var authUser, guest bool
	//user, _ = M.IsGuest(c, G.Store)
	_, _ = M.IsAuthUser(c, G.Store)


	id,_ := strconv.Atoi(c.Param("id"))
	success := S.DeleteFromCart(c, uint(id))
	if success {
		G.Msg.Success = "Product deleted form cart"
		c.Redirect(http.StatusFound, "/show-cart")
	} else {
		G.Msg.Fail = "Some Error Occurred, Please Try Again."
		c.Redirect(http.StatusFound, "/show-cart")
	}
}


