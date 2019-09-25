package Repositories


import (
	Cfg "GoEcommerceProject/Config"

	Mod "GoEcommerceProject/Models"
)

func Orders(orders []Mod.Order, where ...interface{}) []Mod.Order{
	db := Cfg.DBConnect()
	db.Find(&orders, where ...)
	for i, _ := range orders {
		db.Find(&orders[i].User, "id=?", orders[i].UserID)
		db.Find(&orders[i].Bill, "id=?", orders[i].BillID)
		db.Find(&orders[i].PayMethod, "id=?", orders[i].PayMethodID)
		db.Find(&orders[i]).Related(&orders[i].OrderDetails)
	}

	defer db.Close()
	return orders
}


func UpdateOrder(order Mod.Order,values interface{}, where ...interface{}) bool {
	db := Cfg.DBConnect()
	if db.Model(&order).Where(where).Updates(values).RowsAffected == 1 {
		defer db.Close()
		return true
	}
	defer db.Close()
	return false
}


func DeleteOrder(order Mod.Order) bool{
	db := Cfg.DBConnect()
	if db.Delete(&order).RowsAffected == 1 {
		defer db.Close()
		return true
	} else {
		defer db.Close()
		return false
	}
}


func OrderDetails(orderDetails []Mod.OrderDetail) []Mod.OrderDetail {
	db := Cfg.DBConnect()
	for i, _ := range orderDetails {
		db.Find(&orderDetails[i]).Related(&orderDetails[i].Product)
	}

	defer db.Close()
	return orderDetails
}


func UpdateBill(bill Mod.Bill,values interface{}, where ...interface{}) bool{
	db := Cfg.DBConnect()
	if db.Model(&bill).Where(where).Updates(values).RowsAffected == 1 {
		defer db.Close()
		return true
	}
	defer db.Close()
	return false
}