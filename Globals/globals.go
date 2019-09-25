package Globals

import (
	Mod "GoEcommerceProject/Models"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"html/template"
)

type DB_ENV struct {
	Host, Port, Dialect, Username, Password, DBname string
}

type App_env struct {
	Name, Url string
}

type Message struct {
	Success template.HTML
	Fail template.HTML
}

type EmailGenerals struct {
	From, To, Subject, HtmlString string
}




type Cart struct {
	Key string
	Product Mod.Product
	Quantity int
	Total float64
	TotalTax float64
	TotalWithTax float64
}

type FinalCart struct {
	Carts []Cart
	SubTotal float64
	Tax int
	GrandTotalTax float64
	ShippingCost float64
	GrandTotal float64
}

var(
	Store = sessions.NewCookieStore([]byte("secret"))
	DBEnv DB_ENV
	DB *gorm.DB
	//Role Mod.Role
	//User Mod.User
	//PS Mod.PasswordReset
	Msg Message
	AppEnv App_env
)