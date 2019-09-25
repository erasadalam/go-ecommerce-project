package Models

import (
	"bytes"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Order struct{
	gorm.Model
	UserID uint	`gorm:"index; not null"`
	BillID uint	`gorm:"index; not null"`
	PayMethodID uint	`gorm:"index; not null"`
	Total float64	`gorm:"not null"`
	Status int	`gorm:"type:tinyint(4); not null; default:0"`
	User User	`gorm:"save_associations:false; association_save_reference:false"`
	Bill Bill	`gorm:"save_associations:false; association_save_reference:false"`
	PayMethod PayMethod	`gorm:"save_associations:false; association_save_reference:false"`
	OrderDetails []OrderDetail	`gorm:"save_associations:false; association_save_reference:false"`
}

func (u Order) FormatDateTime(t time.Time) string {
	var buffer bytes.Buffer
	//buffer.WriteString(t.Month().String()[:3])
	buffer.WriteString(fmt.Sprintf("%2d - %s - %2d at %2d:%2d", t.Day(), t.Month().String()[:], t.Year(), t.Hour(), t.Minute()))
	return buffer.String()

}