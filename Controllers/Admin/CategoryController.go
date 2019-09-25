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
	Category = make(map[uint]Mod.Category)       //hard to understand
)

func AddCategoryGet(c *gin.Context) {
	if user, success := M.IsAuthAdminUser(c, G.Store); success {
		c.HTML(http.StatusOK, "add-category.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Msg":G.Msg, "Title":"Add-Category"})
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
}

func AddCategoryPost(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var category Mod.Category
	err := c.ShouldBind(&category)         //need to understand
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-category")
		return
	}
	category.Description.String = c.PostForm("description")
	if category.Description.String != "" {
		category.Description.Valid = true
	} else {
		category.Description.Valid = false
	}

	if R.AddCategory(category) {
		G.Msg.Success = "Category Added Successfully."
		c.Redirect(http.StatusFound, "/add-category")
		return
	} else {
		G.Msg.Fail = "Some Error Occured. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-category")
		return
	}
}


func AllCategory(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var categories []Mod.Category
	categories = R.Categories(categories)		//hard to understand
	for _, category := range categories {
		Category[category.ID] = category		//hard to understand
	}
	c.HTML(http.StatusOK, "all-category.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"All-Category", "Categories":categories, "Msg":G.Msg})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}

func MakeCategoryInactive(c *gin.Context) {

	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category = Category[uint(id)]				//hard to understand
	category.Status = 0
	if R.UpdateCategory(category) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-category")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-category")
	}
}


func MakeCategoryActive(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category = Category[uint(id)]
	category.Status = 1
	if R.UpdateCategory(category) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-category")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-category")
	}
}


func EditCategory(c *gin.Context) {
	user, success := M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category = Category[uint(id)]
	c.HTML(http.StatusOK, "edit-category.html",map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"Edit-Category", "Category":category})
}


func UpdateCategory(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.PostForm("id"))
	category = Category[uint(id)]
	category.Name = c.PostForm("name")
	category.Description.String = c.PostForm("description")
	if category.Description.String != "" {
		category.Description.Valid = true
	} else {
		category.Description.Valid = false
	}

	if R.UpdateCategory(category) {
		G.Msg.Success = "Updated Successfully"
		c.Redirect(http.StatusFound, "/all-category")
	} else {
		G.Msg.Success = "Some Error Occurred, Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-category")
	}
}


func DeleteCategory(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var category Mod.Category
	id, _ := strconv.Atoi(c.Param("id"))
	category = Category[uint(id)]
	if R.DeleteCategory(category) {
		G.Msg.Success = "Deleted Successfully"
		c.Redirect(http.StatusFound, "/all-category")
	} else {
		G.Msg.Success = "Some Error Occurred, Deletion Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-category")
	}
}