package application

import (
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type CategoryAppInterface interface {
	GetCategoryPageDetailBySlug(slug string) (*entity.Category, error)
}

type CategoryApp struct {
	CategoryRepo      repository.CategoryRepository
	SectionRepo       repository.SectionRepository
	PublishedPostRepo repository.PublishedPostRepository
}

func (app *CategoryApp) GetCategoryPageDetailBySlug(slug string) (*entity.Category, error) {
	category, err := app.CategoryRepo.GetCategoryBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Category not found
	if category.ID.Int64 == 0 {
		return nil, &common.NotFoundError{}
	}

	ppc, err := app.PublishedPostRepo.GetPublishedPostCountByCategoryId(category.ID.Int64)
	if err != nil {
		return nil, err
	}

	category.PublishedPostCount = ppc

	return category, nil
}
