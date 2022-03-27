package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type CategoryRepository interface {
	GetCategoryBySlug(slug string) (*entity.Category, error)
	GetCategoriesByIds(categoryIds []int64) ([]entity.Category, error)
}
