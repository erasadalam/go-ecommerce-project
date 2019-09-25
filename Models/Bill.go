package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Bill struct{
	gorm.Model
	UserID uint	`gorm:"index; not null" form:"user_id"`
	FullName string	`gorm:"not null" form:"full_name"`
	Email sql.NullString
	Address string `gorm:"not null" form:"address"`
	Phone string	`gorm:"not null" form:"phone"`
	Status int	`gorm:"type:tinyint(4); not null; default:0"`
}