package category

import (
	"product-management-api/pkg/db"
)

type CategoryRepository struct {
	Database *db.DB
}

func NewCategoryRepository(db *db.DB) *CategoryRepository {
	return &CategoryRepository{
		Database: db,
	}
}

func (repo *CategoryRepository) Create(category *Category) (*Category, error) {
	result := repo.Database.DB.Create(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
