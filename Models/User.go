package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	FullName          string         `gorm:"not null" form:"full_name"`
	Email             string         `gorm:"not null; unique_index" form:"email"`
	Phone             string 		 `gorm:"not null; unique_index" form:"phone"`
	PhoneVerification sql.NullString
	Password          string `gorm:"not null"`
	ActiveStatus      int    `gorm:"type:tinyint(4); not null; default:0"`
	RoleID            uint   `gorm:"index; not null"`
	EmailVerification sql.NullString
	RememberToken     sql.NullString
	Role              Role	`gorm:"save_associations:false; association_save_reference:false"`
	Orders            []Order	`gorm:"save_associations:false; association_save_reference:false"`
	Bills             []Bill	`gorm:"save_associations:false; association_save_reference:false"`
	Wishlist          Wishlist	`gorm:"save_associations:false; association_save_reference:false"`
}

