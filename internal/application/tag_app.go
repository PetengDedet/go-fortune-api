package application

import (
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type TagAppInterface interface {
	GetTagPageDetail(tagSlug string) (*entity.Page, error)
}

type TagApp struct {
	TagRepo           repository.TagRepository
	SectionRepo       repository.SectionRepository
	PublishedPostRepo repository.PublishedPostRepository
}

func (app *TagApp) GetTagDetail(tagSlug string) (*entity.Tag, error) {
	tag, err := app.TagRepo.GetTagBySlug(tagSlug)
	if err != nil {
		panic(err)
	}

	// Tag not found
	if tag.ID == 0 {
		return nil, &common.NotFoundError{}
	}

	ppc, err := app.PublishedPostRepo.GetPublishedPostCountByTagId(tag.ID)
	if err != nil {
		return nil, err
	}

	tag.PublishedPostCount = ppc

	return tag, nil
}
