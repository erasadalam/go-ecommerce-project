package Admin

import (
	"GoEcommerceProject/Controllers/Admin"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	r.GET("/dashboard", Admin.Dashboard)

	r.GET("/add-brand", Admin.AddBrandGet)
	r.POST("/add-brand", Admin.AddBrandPost)
	r.GET("/all-brand", Admin.AllBrand)
	r.GET("/make-brand-inactive/:id", Admin.MakeBrandInactive)
	r.GET("/make-brand-active/:id", Admin.MakeBrandActive)
	r.GET("/edit-brand/:id", Admin.EditBrand)
	r.POST("/update-brand", Admin.UpdateBrand)
	r.GET("/delete-brand/:id", Admin.DeleteBrand)

	r.GET("/add-category", Admin.AddCategoryGet)
	r.POST("/add-category", Admin.AddCategoryPost)
	r.GET("/all-category", Admin.AllCategory)

	r.GET("/make-category-inactive/:id", Admin.MakeCategoryInactive)
	r.GET("/make-category-active/:id", Admin.MakeCategoryActive)
	r.GET("/edit-category/:id", Admin.EditCategory)
	r.POST("/update-category", Admin.UpdateCategory)
	r.GET("/delete-category/:id", Admin.DeleteCategory)

	r.GET("/add-product", Admin.AddProductGet)
	r.POST("/add-product", Admin.AddProductPost)
	r.GET("/all-product", Admin.AllProduct)

	r.GET("/make-product-inactive/:id", Admin.MakeProductInactive)
	r.GET("/make-product-active/:id", Admin.MakeProductActive)
	r.GET("/edit-product/:id", Admin.EditProduct)
	r.POST("/update-product", Admin.UpdateProduct)
	r.GET("/delete-product/:id", Admin.DeleteProduct)

	r.GET("/orders", Admin.Orders)
	r.GET("/make-order-pending/:id", Admin.MakeOrderPending)
	r.GET("/make-order-delivered/:id", Admin.MakeOrderDelivered)
	r.GET("/order-details/:id", Admin.OrderDetails)
	r.GET("/delete-order/:id", Admin.DeleteOrder)
	r.GET("/make-payment-pending/:id", Admin.MakePaymentPending)
	r.GET("/make-payment-done/:id", Admin.MakePaymentDone)


}
