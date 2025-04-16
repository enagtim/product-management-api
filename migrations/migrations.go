package migrations

import (
	"product-management-api/internal/category"
	"product-management-api/internal/product"
	"product-management-api/pkg/db"
)

func Migrate(db *db.DB) {
	err := db.AutoMigrate(&category.Category{}, &product.Product{})
	if err != nil {
		panic(err)
	}
}
