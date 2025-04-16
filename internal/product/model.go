package product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	CategoryID  uint64  `gorm:"index:category_id;not null"`
	Price       float64 `gorm:"decimal(10,2);index:price;not null"`
	Quantity    uint64  `gorm:"default:1;not null;"`
	Name        string  `gopm:"index:name;not null"`
	IsAvailable bool    `gorm:"default:true;not null"`
}
