package repository

import "github.com/PetengDedet/fortune-post-api/domain/entity"

type CategoryRepository interface {
	// GetCategoryPageBySlug(slug string) (*entity.CategoryPage, error)
	GetCategoriesByIds(categoryIds []int64) ([]entity.Category, error)
}
