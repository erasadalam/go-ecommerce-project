package Admin

import (
	G "GoEcommerceProject/Globals"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


var (
	Order = make(map[uint]Mod.Order)
)

func Orders(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var orders []Mod.Order
	orders = R.Orders(orders)



	for _, order := range orders {
		Order[order.ID] = order
	}

	date :=  Mod.Order{

	}

	c.HTML(http.StatusOK, "orders.html", map[string]interface{}{
		"AppEnv": G.AppEnv, "User": user, "Title": "Orders", "Msg": G.Msg, "Orders": orders, "Date": date })
	G.Msg.Success = ""
	G.Msg.Fail = ""
}

func MakeOrderPending(c *gin.Context) {
	_, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var order Mod.Order
	order.ID = uint(id)
	if R.UpdateOrder(order, map[string]interface{}{"status": 0}, "id=?", order.ID) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/orders")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/orders")
	}

}


func MakeOrderDelivered(c *gin.Context) {
	_, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var order Mod.Order
	order.ID = uint(id)
	if R.UpdateOrder(order, map[string]interface{}{"status": 1}, "id=?", order.ID) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/orders")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/orders")
	}

}


func OrderDetails(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var order Mod.Order
	id, _ := strconv.Atoi(c.Param("id"))
	order = Order[uint(id)]
	order.OrderDetails = R.OrderDetails(order.OrderDetails)
	c.HTML(http.StatusOK, "order-details.html",map[string]interface{}{
		"AppEnv":G.AppEnv, "User":user, "Title":"Order-Details", "Order":order, "Msg":G.Msg})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}


func DeleteOrder(c *gin.Context) {
	_, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var order Mod.Order
	order.ID = uint(id)
	if R.DeleteOrder(order) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/orders")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/orders")
	}
}

func MakePaymentPending(c *gin.Context) {
	_, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var bill Mod.Bill
	bill.ID = uint(id)
	if R.UpdateBill(bill, map[string]interface{}{"status": 0}, "id=?", bill.ID) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/orders")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/orders")
	}
}


func MakePaymentDone(c *gin.Context) {
	_, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var bill Mod.Bill
	bill.ID = uint(id)
	if R.UpdateBill(bill, map[string]interface{}{"status": 1}, "id=?", bill.ID) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/orders")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/orders")
	}
}