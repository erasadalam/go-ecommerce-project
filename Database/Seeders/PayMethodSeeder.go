package Seeders

import (
	Cfg "GoEcommerceProject/Config"

	Mod "GoEcommerceProject/Models"
)

var payMethods = make([]Mod.PayMethod,0)

func PayMethodSeeder() {

	payMethod1()
	payMethod2()
	payMethod3()
	db := Cfg.DBConnect()
	for i,_ := range payMethods {
		db.FirstOrCreate(&payMethods[i],&Mod.PayMethod{Method:payMethods[i].Method})
	}
}

func payMethod1() {
	var payMethod = Mod.PayMethod{
		Method: "Cash on Delivery",
		Status: 1,
	}
	payMethods = append(payMethods, payMethod)
}

func payMethod2() {
	var payMethod = Mod.PayMethod{
		Method: "bKash",
		Status: 1,
	}
	payMethods = append(payMethods, payMethod)
}


func payMethod3() {
	var payMethod = Mod.PayMethod{
		Method: "PayPal",
		Status: 1,
	}
	payMethods = append(payMethods, payMethod)
}
