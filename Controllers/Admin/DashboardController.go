package Admin

import (
	G "GoEcommerceProject/Globals"
	M "GoEcommerceProject/Middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dashboard(c *gin.Context) {

	if user, success := M.IsAuthAdminUser(c, G.Store); success {
		c.HTML(http.StatusOK, "dashboard.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"Dashboard"})
	}
	return
}
