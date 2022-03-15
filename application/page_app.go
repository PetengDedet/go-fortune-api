package application

import (
	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type PageAppInterface interface {
	GetPageDetailBySlug(slug string) (*entity.Page, error)
}
type PageApp struct {
	PageRepo    repository.PageRepository
	SectionRepo repository.SectionRepository
}

// var _ PageAppInterface = &PageApp{}

func (pageApp *PageApp) GetPageDetailBySlug(slug string) (*entity.Page, error) {
	page, err := pageApp.PageRepo.GetPageBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Empty page
	if page.ID == 0 {
		return nil, &common.NotFoundError{}
	}

	sections, err := pageApp.SectionRepo.GetSectionsByPageId(page.ID)
	if err != nil {
		return nil, err
	}

	var ss []entity.Section
	for _, s := range sections {
		ss = append(ss, *entity.SectionResponse(&s))
	}
	page.Sections = ss

	if len(ss) <= 0 {
		page.Sections = []entity.Section{}
	}

	// return entity.PageResponse(page), nil
	return nil, nil
}
