package Auth

import (
	G "GoEcommerceProject/Globals"
	H "GoEcommerceProject/Helpers"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

func LoginGet(c *gin.Context) {

	if user, success := M.IsGuest(c, G.Store); success {
		session, _ := G.Store.Get(c.Request, "cart")
		cart_items := len(session.Values)
		request_url := c.Query("q")
		c.HTML(http.StatusOK, "login.html", map[string]interface{}{
			"AppEnv": G.AppEnv, "User": user, "Msg": G.Msg, "Title": "Login", "Cart_items": cart_items, "Request_url": request_url})
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
	return
}


/*func RegisterGet(c *gin.Context) {

	if _, success := M.IsGuest(c, G.Store); success {
		c.HTML(http.StatusOK, "register.html", G.Msg)
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
	return
}*/

func RegisterPost(c *gin.Context) {
	var success bool
	var user Mod.User
	user.FullName = c.PostForm("full_name")
	user.Email = c.PostForm("email")
	_, success = R.ReadWithEmail(user)
	if success {
		G.Msg.Fail = "User Already Exists With This Email."
		c.Redirect(http.StatusFound, "/login")
		return
	}
	user.Phone = c.PostForm("phone")
	_, success = R.ReadWithPhone(user)
	if success {
		G.Msg.Fail = "User Already Exists With This Phone Number."
		c.Redirect(http.StatusFound, "/login")
		return
	}
	password := c.PostForm("password")
	confirmPass := c.PostForm("confirm-password")
	if password != confirmPass {
		G.Msg.Fail = "Confirm Password Doesn't Match."
		c.Redirect(http.StatusFound, "/login")
		return
	}
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), cost)
	user.Password = string(hash)
	user.RoleID = 2
	user.ActiveStatus = 1
	user.RememberToken.String = H.RandomString(60)
	user.RememberToken.Valid = true
	user, success = R.Register(user)
	if success {
		session, _ := G.Store.Get(c.Request, "login_token")
		session.Values["userEmail"] = user.Email
		session.Values["remember_token"] = user.RememberToken.String
		session.Save(c.Request, c.Writer)

		session, _ = G.Store.Get(c.Request, "cart")
		if len(session.Values) > 0 {
			c.Redirect(http.StatusFound, "/show-cart")
		}else {
			c.Redirect(http.StatusFound, "/")
		}
	} else {
		if G.Msg.Fail == "" {
			G.Msg.Fail = "Some Internal Server Error Occurred, Please Try Again."
		}
		c.Redirect(http.StatusFound, "/login")
	}
}

func LoginPost(c *gin.Context) {

	var user Mod.User
	user.Email = c.PostForm("email")
	password := c.PostForm("password")
	//rememberMe, _ := strconv.Atoi(c.PostForm("remember_me"))
	var success bool
	user, success = R.Login(user)
	if success {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			G.Msg.Fail = "Wrong Credentials."
			c.Redirect(http.StatusFound, "/login")
		} else {
			if user.ActiveStatus == 1 {
				user.RememberToken.String = H.RandomString(60)
				user.RememberToken.Valid = true

				if !R.SetRememberToken(user) {
					G.Msg.Fail = "Some Internal Server Error Occurred. Please Try Again."
					c.Redirect(http.StatusFound, "/login")
					return
				}

				session, _ := G.Store.Get(c.Request, "login_token")
				session.Values["userEmail"] = user.Email
				session.Values["remember_token"] = user.RememberToken.String
				//session.Options.MaxAge = 60 * 60 * 24 * 5
				session.Save(c.Request, c.Writer)

				/*if rememberMe == 1 {
					session.Options.MaxAge = 60 * 60 * 24 * 365
					session.Save(c.Request, c.Writer)
				}*/

				if user.RoleID == 1 {
					c.Redirect(http.StatusFound, "/dashboard")
				} else if user.RoleID == 2 {
					session.Save(c.Request, c.Writer)

					var wishlists []Mod.Wishlist
					wishlists = R.WishlistWithProducts(wishlists, user.ID)
					session.Values["wish_items"] = len(wishlists)
					session.Save(c.Request, c.Writer)

					request_url := c.Query("q")

					session, _ := G.Store.Get(c.Request, "cart")
					if request_url != "" {
						c.Redirect(http.StatusFound, request_url)
					} else if len(session.Values) > 0 {
						c.Redirect(http.StatusFound, "/show-cart")
					} else {
						c.Redirect(http.StatusFound, "/")
					}
				}
			} else if user.ActiveStatus == 2 {
				if G.Msg.Fail == "" {
					G.Msg.Fail = "You Are Suspended. Contact With The Authority Quickly."
				}
				c.Redirect(http.StatusFound, "/login")
			} else {
				var link template.HTML
				link = "<a href='http://localhost:8080/resend-email-verification'>Click Here To Send Verification Email</a>"
				if G.Msg.Fail == "" {
					G.Msg.Fail = "Please Activate Your Account, " + link + "."
				}

				c.Redirect(http.StatusFound, "/")
			}
		}

	} else {
		if G.Msg.Fail == "" {
			G.Msg.Fail = "User Not Found."
		}
		c.Redirect(http.StatusFound, "/login")
	}
}

func Logout(c *gin.Context) {
	var user Mod.User

	session, _ := G.Store.Get(c.Request, "login_token")
	user.Email = session.Values["userEmail"].(string)
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)

	R.Logout(user)

	c.Redirect(http.StatusFound, "/")
}
