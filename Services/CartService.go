package Services

import (
	Cfg "GoEcommerceProject/Config"
	G "GoEcommerceProject/Globals"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	"github.com/gin-gonic/gin"
	"log"
)

func AddCart(c * gin.Context, id int, quantity int) bool {
	session, _ := G.Store.Get(c.Request, "cart")

	var product Mod.Product
	product.ID = uint(id)
		if session.Values[product.ID] == nil {
			session.Values[product.ID] = quantity
			G.Msg.Success = "Product added into cart"
		} else if (session.Values[product.ID].(int) + quantity) > 20 {
			G.Msg.Fail = "You Have exceed the limit of products adding. Max limit 20"
			session.Values[product.ID] = 20
		} else {
			session.Values[product.ID] = session.Values[product.ID].(int) + quantity
			G.Msg.Success = "Product added into cart"
		}


	/*flag := 0
	if len(session.Values) > 0 {
		for key,value := range session.Values {
			//value := session.Values[key]
			if key == product.ID {
				quantity = quantity + value.(int)
				if quantity > 10 {
					G.Msg.Fail = "You Have exced the limit of products adding. Max limit 10"
					quantity = 10
				}else{
					session.Values[key] = quantity
					flag = 1
					break
				}

			}
		}
		if flag == 0 {
			session.Values[product.ID] = quantity
		}
	}else {
		session.Values[product.ID] = quantity
	}*/

	err := session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println("CartService log1:", err.Error())
		return false
	}
	return true
}

func ProcessCart(c *gin.Context) G.FinalCart {
	session, _ := G.Store.Get(c.Request, "cart")
	var carts []G.Cart

	var Total float64
	var TotalTax float64

	var FinalSubTotal float64
	var FinalTotalTax float64

	for key,value := range session.Values {
		var product Mod.Product
		product.ID = key.(uint)
		product = R.Product(product)
		var cart G.Cart
		cart.Key = ""
		cart.Product = product
		cart.Quantity = value.(int)
		TotalTax = float64(cart.Quantity) * (product.Price * Cfg.CartTax / 100.0)
		cart.TotalTax = TotalTax
		FinalTotalTax += TotalTax
		Total =  float64(cart.Quantity) * product.Price
		cart.Total = Total
		FinalSubTotal += Total
		cart.TotalWithTax = Total + TotalTax
		carts = append(carts, cart)
	}
	var finalCart G.FinalCart
	finalCart.Carts = carts
	finalCart.SubTotal = FinalSubTotal
	finalCart.Tax = int(Cfg.CartTax)
	finalCart.GrandTotalTax = FinalTotalTax
	finalCart.ShippingCost = Cfg.ShippingCost
	finalCart.GrandTotal = FinalSubTotal + FinalTotalTax + Cfg.ShippingCost

	return finalCart
}


func DestroyCart(c *gin.Context) bool{
	session, _ := G.Store.Get(c.Request, "cart")
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)

	if err != nil {
		log.Println("CartService log2:", err.Error())
		return false
	}
	return true

}


func DeleteFromCart(c *gin.Context, id uint) bool{
	session, _ := G.Store.Get(c.Request, "cart")
	delete(session.Values, id)
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println("CartService log3:", err.Error())
		return false
	}
	return true
}

