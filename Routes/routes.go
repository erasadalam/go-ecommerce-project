package Routes

import (
	"GoEcommerceProject/Routes/Admin"
	"GoEcommerceProject/Routes/Auth"
	"GoEcommerceProject/Routes/User"
	"github.com/gin-gonic/gin"
)

func Routes() {

	r := gin.Default()
	r.Static("/assets", "./")
	r.LoadHTMLGlob("Views/**/*.html")

	Auth.AuthRoutes(r)
	Admin.AdminRoutes(r)
	User.UserRoutes(r)

	r.Run(":8080")
}
