package Admin

import (
	G "GoEcommerceProject/Globals"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

var (
	Brand = make(map[uint]Mod.Brand)
)

func AddBrandGet(c *gin.Context) {
	if user, success := M.IsAuthAdminUser(c, G.Store); success {
		c.HTML(http.StatusOK, "add-brand.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Msg":G.Msg, "Title":"Add-Brand"})
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
}

func AddBrandPost(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var brand Mod.Brand
	err := c.ShouldBind(&brand)
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-brand")
		return
	}
	brand.Description.String = c.PostForm("description")
	if brand.Description.String != "" {
		brand.Description.Valid = true
	} else {
		brand.Description.Valid = false
	}
	if R.AddBrand(brand) {
		G.Msg.Success = "Brand Added Successfully."
		c.Redirect(http.StatusFound, "/add-brand")
		return
	} else {
		G.Msg.Fail = "Some Error Occured. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-brand")
		return
	}
}


func AllBrand(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var brands []Mod.Brand
	brands = R.Brands(brands)
	for _, brand := range brands {
		Brand[brand.ID] = brand
	}
	c.HTML(http.StatusOK, "all-brand.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"All-Brand", "Brands":brands, "Msg":G.Msg})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}

func MakeBrandInactive(c *gin.Context) {

	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.Param("id"))
	brand = Brand[uint(id)]
	brand.Status = 0
	if R.UpdateBrand(brand) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-brand")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-brand")
	}
}


func MakeBrandActive(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.Param("id"))
	brand = Brand[uint(id)]
	brand.Status = 1
	if R.UpdateBrand(brand) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-brand")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-brand")
	}
}


func EditBrand(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.Param("id"))
	brand = Brand[uint(id)]
	c.HTML(http.StatusOK, "edit-brand.html",map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"Edit-Brand", "Brand":brand})
}


func UpdateBrand(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.PostForm("id"))
	brand = Brand[uint(id)]
	brand.Name = c.PostForm("name")
	brand.Description.String = c.PostForm("description")
	if brand.Description.String != "" {
		brand.Description.Valid = true
	} else {
		brand.Description.Valid = false
	}

	if R.UpdateBrand(brand) {
		G.Msg.Success = "Updated Successfully"
		c.Redirect(http.StatusFound, "/all-brand")
	} else {
		G.Msg.Success = "Some Error Occurred, Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-brand")
	}
}


func DeleteBrand(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var brand Mod.Brand
	id, _ := strconv.Atoi(c.Param("id"))
	brand = Brand[uint(id)]
	if R.DeleteBrand(brand) {
		G.Msg.Success = "Deleted Successfully"
		c.Redirect(http.StatusFound, "/all-brand")
	} else {
		G.Msg.Success = "Some Error Occurred, Deletion Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-brand")
	}
}