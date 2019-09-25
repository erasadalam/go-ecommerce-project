package Admin

import (
	G "GoEcommerceProject/Globals"
	H "GoEcommerceProject/Helpers"
	M "GoEcommerceProject/Middlewares"
	Mod "GoEcommerceProject/Models"
	R "GoEcommerceProject/Repositories"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

var (
	Product = make(map[uint]Mod.Product)
)

func AddProductGet(c *gin.Context) {
	if user, success := M.IsAuthAdminUser(c, G.Store); success {
		var categories []Mod.Category
		var brands []Mod.Brand

		var products []Mod.Product
		products = R.ProductsWithOthers(products)


		categories = R.Categories(categories, "status=?", 1)     //got hard to understand
		brands = R.Brands(brands, "status=?", 1)
		c.HTML(http.StatusOK, "add-product.html", map[string]interface{}{
			"AppEnv":G.AppEnv, "User":user, "Msg":G.Msg, "Title":"Add-Product", "Categories":categories, "Brands":brands, "Product_positions": Iterate(uint(len(products) + 1))})
		G.Msg.Success = ""
		G.Msg.Fail = ""
	}
}

func Iterate(count uint) []uint {
	var i uint
	var Items []uint
	for i = 1; i < (count+1); i++ {
		Items = append(Items, i)
	}
	return Items
}

func AddProductPost(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var product Mod.Product
	err := c.ShouldBind(&product)   //hard to understand it
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Some Error Occurred. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-product")
		return
	}
	product.Description.String = c.PostForm("description")
	product.Description = H.NullStringProcess(product.Description)
	product.Size.String = c.PostForm("size")
	product.Size = H.NullStringProcess(product.Size)
	product.Color.String = c.PostForm("color")
	product.Color = H.NullStringProcess(product.Color)
/*	pp,_ := strconv.Atoi(c.PostForm("product_sl"))
	product.ProductSL = uint(pp)*/
	img, _ := c.FormFile("img")			//need to understand
	ext := filepath.Ext(img.Filename)			//
	imgName := H.RandomString(60)+ext
	dst := "./Storage/Images/"+imgName
	err = c.SaveUploadedFile(img, dst)			//need to understand
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Image Upload Failed. Try Again Later."
		c.Redirect(http.StatusFound, "/add-product")
		return
	}
	imgUrl := []byte(dst)
	product.ImgUrl.String = string(imgUrl[1:])			//need to understand
	product.ImgUrl = H.NullStringProcess(product.ImgUrl)

	if R.AddProduct(product) {
		G.Msg.Success = "Product Added Successfully."
		c.Redirect(http.StatusFound, "/add-product")
		return
	} else {
		G.Msg.Fail = "Some Error Occured. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-product")
		return
	}
}


func AllProduct(c *gin.Context) {
	var user Mod.User
	var success bool
	user, success = M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var products []Mod.Product
	products = R.ProductsWithOthers(products)
	for _, product := range products {
		Product[product.ID] = product
	}
	c.HTML(http.StatusOK, "all-product.html", map[string]interface{}{"AppEnv":G.AppEnv, "User":user, "Title":"All-Product", "Products":products, "Msg":G.Msg})
	G.Msg.Success = ""
	G.Msg.Fail = ""
}

func MakeProductInactive(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = Product[uint(id)]
	product.Status = 0
	if R.UpdateProduct(product) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-product")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-product")
	}
}


func MakeProductActive(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = Product[uint(id)]
	product.Status = 1
	if R.UpdateProduct(product) {
		G.Msg.Success = "Status Updated Successfully"
		c.Redirect(http.StatusFound, "/all-product")
	} else {
		G.Msg.Success = "Some Error Occurred, Status Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-product")
	}
}


func EditProduct(c *gin.Context) {
	var user Mod.User
	var success bool
	user, success = M.IsAuthAdminUser(c, G.Store)
	if !success {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = Product[uint(id)]
	var categories []Mod.Category
	var brands []Mod.Brand

	var products []Mod.Product
	products = R.ProductsWithOthers(products)

	categories = R.Categories(categories, "status=? and id <> ?", 1, product.CategoryID)
	brands = R.Brands(brands, "status=? and id <> ?", 1, product.BrandID)
		c.HTML(http.StatusOK, "edit-product.html",map[string]interface{}{
		"AppEnv":G.AppEnv, "User":user, "Title":"Edit-Product", "Product":product, "Categories":categories, "Brands":brands, "Product_positions": Iterate(uint(len(products)))})
}


func UpdateProduct(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.PostForm("id"))
	product = Product[uint(id)]
	err := c.ShouldBind(&product)
	if err != nil {
		log.Println(err.Error())
		G.Msg.Fail = "Some Error Occured. Please Try Again Later."
		c.Redirect(http.StatusFound, "/add-product")
		return
	}
	product.Description.String = c.PostForm("description")
	product.Description = H.NullStringProcess(product.Description)
	product.Size.String = c.PostForm("size")
	product.Size = H.NullStringProcess(product.Size)
	product.Color.String = c.PostForm("color")
	product.Color = H.NullStringProcess(product.Color)
	img, _ := c.FormFile("img")
	if img != nil {

		os.Remove("."+product.ImgUrl.String)

		ext := filepath.Ext(img.Filename)
		imgName := H.RandomString(60) + ext
		dst := "./Storage/Images/" + imgName
		err = c.SaveUploadedFile(img, dst)
		if err != nil {
			log.Println(err.Error())
			G.Msg.Fail = "Image Upload Failed. Try Again Later."
			c.Redirect(http.StatusFound, "/add-product")
			return
		}
		imgUrl := []byte(dst)
		product.ImgUrl.String = string(imgUrl[1:])
		product.ImgUrl = H.NullStringProcess(product.ImgUrl)
	}

	if R.UpdateProduct(product) {
		G.Msg.Success = "Updated Successfully"
		c.Redirect(http.StatusFound, "/all-product")
	} else {
		G.Msg.Success = "Some Error Occurred, Update Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-product")
	}
}


func DeleteProduct(c *gin.Context) {
	if _, success := M.IsAuthAdminUser(c, G.Store); !success {
		return
	}
	var product Mod.Product
	id, _ := strconv.Atoi(c.Param("id"))
	product = Product[uint(id)]
	if R.DeleteProduct(product) {
		G.Msg.Success = "Deleted Successfully"
		c.Redirect(http.StatusFound, "/all-product")
	} else {
		G.Msg.Success = "Some Error Occurred, Deletion Failed. Please Try Again Later."
		c.Redirect(http.StatusFound, "/all-product")
	}
}