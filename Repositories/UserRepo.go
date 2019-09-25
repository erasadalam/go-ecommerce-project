package Repositories

import (
	Cfg "GoEcommerceProject/Config"
	Mod "GoEcommerceProject/Models"
)

func PayMethods(payMethods []Mod.PayMethod, where ...interface{}) []Mod.PayMethod{
	db := Cfg.DBConnect()
	db.Find(&payMethods, where...)
	defer db.Close()
	return payMethods
}


func AddBill(bill Mod.Bill) (Mod.Bill, bool) {
	db := Cfg.DBConnect()
	db.Create(&bill)
	if bill.ID != 0 {
		return bill, true
	}
	defer db.Close()
	return bill, false
}


func AddOrder(order Mod.Order) (Mod.Order, bool){
	db := Cfg.DBConnect()
	db.Create(&order)
	if order.ID !=0 {
		return order, true
	}
	defer db.Close()
	return order, false
}


func AddOrderDetails(orderDetails []Mod.OrderDetail) ([]Mod.OrderDetail, bool){
	db := Cfg.DBConnect()
	for i, _ := range orderDetails {
		db.Create(&orderDetails[i])
		if orderDetails[i].ID == 0 {
			return orderDetails, false
		}
	}
	defer db.Close()
	return orderDetails, true
}