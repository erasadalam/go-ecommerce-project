package User

import (
	G "GoEcommerceProject/Globals"
	H "GoEcommerceProject/Helpers"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	S "GoEcommerceProject/Services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)


var (
	Product = make(map[uint]Mod.Product)
)

func Welcome(c *gin.Context) {
	session, _ := G.Store.Get(c.Request, "cart")
	var user Mod.User
	//var authUser, guest bool
	//user, _ = M.IsGuest(c, G.Store)
	user, _ = M.IsAuthUser(c, G.Store)
	//user, _ = M.IsAuthAdminUser(c, G.Store)
/*	if !authUser && !guest {
		return
	}*/
	var categories []Mod.Category
	var brands []Mod.Brand
	var products []Mod.Product
	categories = R.Categories(categories, "status=?", 1)
	brands = R.Brands(brands, "status=?", 1)
	products = R.ProductsWithOthers(products, "status=?", 1)
	for _, product := range products {		//for each product(row) from products, struct type variable
		Product[product.ID] = product		//we insert each row in Product, map array, with respect to productID
	}

		var wishlists []Mod.Wishlist
		wishlists = R.WishlistWithProducts(wishlists, user.ID)
		session_wishlist, _ := G.Store.Get(c.Request, "login_token")
		session_wishlist.Values["wish_items"] = len(wishlists)
		session_wishlist.Save(c.Request, c.Writer)
		//"Wish_items": session_wishlist.Values["wish_items"]




	cart_items := len(session.Values)
	c.HTML(http.StatusOK, "welcome.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Welcome", "Categories": categories, "Brands": brands, "Products": products, "Cart_items": cart_items, "Wish_items": session_wishlist.Values["wish_items"]})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}

func ProductDetails(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = Product[uint(id)]			//from the map type variable Product, we fetch a row of previously
										// inserted of productID
	session, _ := G.Store.Get(c.Request, "cart")
	cart_items := len(session.Values)

	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)
	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]

	c.HTML(http.StatusOK, "product-details.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Product-Details", "Product": product, "Cart_items": cart_items,"Wish_items": session_wishlist.Values["wish_items"]})
}


func BrandWisePro(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.Param("id"))
	brand.ID = uint(id)



	brand = R.BrandWithProducts(brand)


	session, _ := G.Store.Get(c.Request, "cart")
	cart_items := len(session.Values)

	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)
	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]

	c.HTML(http.StatusOK, "brand-wise-pro.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Brand-Wise-Product", "Brand": brand, "Cart_items": cart_items,"Wish_items": session_wishlist.Values["wish_items"] })
}


func CatWisePro(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = uint(id)
	category = R.CategoryWithProducts(category)
	fmt.Println(category)
	session, _ := G.Store.Get(c.Request, "cart")
	cart_items := len(session.Values)

	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)
	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]

	c.HTML(http.StatusOK, "cat-wise-pro.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Category-Wise-Product", "Category": category, "Cart_items": cart_items,"Wish_items": session_wishlist.Values["wish_items"]})
}



func Checkout(c *gin.Context) {
	var authUser, guest bool
	_, guest = M.IsGuest(c, G.Store)
	_, authUser = M.IsAuthUser(c, G.Store)

	total,_ := strconv.Atoi(c.Param("total"))
	if authUser && total > 0{
		c.Redirect(http.StatusFound, "/bill-to")
	} else if guest && total > 0 {
		c.Redirect(http.StatusFound, "/login")
	} else {
		c.Redirect(http.StatusFound, "/")
	}
}


func BillToGet(c *gin.Context) {
	var user Mod.User
	var authUser, guest bool
	user, guest = M.IsGuest(c, G.Store)
	user, authUser = M.IsAuthUser(c, G.Store)
	if !authUser && !guest {
		return
	}
	var payMethods []Mod.PayMethod
	payMethods = R.PayMethods(payMethods,"status=1")

	var wishlists []Mod.Wishlist
	wishlists = R.WishlistWithProducts(wishlists, user.ID)
	session_wishlist, _ := G.Store.Get(c.Request, "login_token")
	session_wishlist.Values["wish_items"] = len(wishlists)
	session_wishlist.Save(c.Request, c.Writer)
	//"Wish_items": session_wishlist.Values["wish_items"]


	c.HTML(http.StatusOK, "bill-to.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Category-Bill-To", "PayMethods":payMethods, "Wish_items": session_wishlist.Values["wish_items"] })
}


func BillToPost(c *gin.Context) {
	var success bool
	var bill Mod.Bill

	err := c.ShouldBind(&bill)
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/bill-to")
		return
	}
	bill.Email.String = c.PostForm("email")
	bill.Email = H.NullStringProcess(bill.Email)
	bill, success = R.AddBill(bill)
	if !success {
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/bill-to")
		return
	}

	finalCart := S.ProcessCart(c)
	var order Mod.Order
	payMethodID, _ := strconv.Atoi(c.PostForm("pay_method_id"))
	order.ID = bill.ID
	order.UserID = bill.UserID
	order.BillID = bill.ID
	order.PayMethodID =  uint(payMethodID)
	order.Total = finalCart.GrandTotal
	order, success = R.AddOrder(order)
	if !success {
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/bill-to")
		return
	}

	var orderDetail Mod.OrderDetail
	var orderDetails []Mod.OrderDetail
	for _, cart := range finalCart.Carts {
		orderDetail.OrderID = order.ID
		orderDetail.ProductID = cart.Product.ID
		orderDetail.Quantity = cart.Quantity
		orderDetail.Total = cart.Total
		orderDetail.TotalTax = cart.TotalTax
		orderDetail.TotalWithTax = cart.TotalWithTax
		orderDetails = append(orderDetails, orderDetail)
	}
	orderDetails, success = R.AddOrderDetails(orderDetails)
	if !success {
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/bill-to")
		return
	}
	S.DestroyCart(c)
	G.Msg.Success = `You have made your order successfully. We will communicate with your billing contact
	number within 30 minutes. Thank you.`
	c.Redirect(http.StatusFound, "/")
}

















