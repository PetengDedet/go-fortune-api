package application

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type CategoryAppInterface interface {
	GetCategoryPageDetailBySlug(slug string) (*entity.CategoryPage, error)
}

type CategoryApp struct {
	CategoryRepo repository.CategoryRepository
	SectionRepo  repository.SectionRepository
}

func (categoryApp *CategoryApp) GetCategoryPageDetailBySlug(slug string) (*entity.CategoryPage, error) {
	category, err := categoryApp.CategoryRepo.GetCategoryPageBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Category not found
	if category.ID == 0 {
		return nil, &common.NotFoundError{}
	}

	sections, err := categoryApp.SectionRepo.GetSectionsByPageSlug("category")
	if err != nil {
		return nil, err
	}

	log.Println("Sectionsss", sections)

	var ss []entity.Section
	for _, s := range sections {
		ss = append(ss, *entity.SectionResponse(&s))
	}
	category.Sections = ss

	if len(ss) <= 0 {
		category.Sections = []entity.Section{}
	}

	return category, nil
}
