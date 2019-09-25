package Auth

import (
	"GoEcommerceProject/Controllers/Auth"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	//r.GET("/register", Auth.RegisterGet)
	r.POST("/register", Auth.RegisterPost)
	r.GET("/login", Auth.LoginGet)
	r.POST("/login", Auth.LoginPost)
	r.GET("/logout", Auth.Logout)
}
