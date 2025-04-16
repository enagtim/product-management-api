package main

import (
	"fmt"
	"net/http"
	"product-management-api/configs"
	"product-management-api/internal/category"
	"product-management-api/migrations"
	"product-management-api/pkg/db"
)

func main() {

	conf := configs.LoadConfig()
	router := http.NewServeMux()
	db := db.NewDb(conf)
	migrations.Migrate(db)

	categoryRepository := category.NewCategoryRepository(db)

	categoryService := category.NewCategoryService(categoryRepository)

	category.NewCategoryHandler(router, categoryService)

	server := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	fmt.Println("Server start on port 8000")
	server.ListenAndServe()
}
