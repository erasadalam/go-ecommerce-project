package Models

import (
	"github.com/jinzhu/gorm"
)

type OrderDetail struct{
	gorm.Model
	OrderID uint	`gorm:"index; not null"`
	ProductID uint	`gorm:"index; not null"`
	Quantity int	`gorm:"not null"`
	Total float64	`gorm:"not null"`
	TotalTax float64	`gorm:"not null"`
	TotalWithTax float64	`gorm:"not null"`
	Product Product	`gorm:"save_associations:false; association_save_reference:false"`
}