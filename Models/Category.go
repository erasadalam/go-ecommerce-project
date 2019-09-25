package Models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Category struct{
	gorm.Model
	Name string	`gorm:"not null" form:"name"`
	Description sql.NullString
	Status int	`gorm:"type:tinyint(4); not null; default:0" form:"status"`

	// Disable auto update/create associations, will update reference for records have primary key.
	//Don't even save association's reference when updating/saving data.
	Products []Product	`gorm:"save_associations:false; association_save_reference:false"`
}
