package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Product struct{
	gorm.Model
	CategoryID uint	`gorm:"index; not null" form:"category_id"`
	BrandID uint	`gorm:"index; not null" form:"brand_id"`

	Name string	`gorm:"not null" form:"name"`
	Price float64	`gorm:"not null" form:"price"`
	Size sql.NullString
	Color sql.NullString
	Description sql.NullString

	ProductSL uint `gorm:"index; not null" form:"product_sl"`

	Status int	`gorm:"type:tinyint(4); not null; default:0" form:"status"`
	ImgUrl sql.NullString
	Category Category	`gorm:"save_associations:false; association_save_reference:false"`
	Brand Brand	`gorm:"save_associations:false; association_save_reference:false"`

}