package application

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type PostTypeAppInterface interface {
	GetPostTypeDetail(slug string) (*entity.PostType, error)
}

type PostTypeApp struct {
	PostTypeRepo      repository.PostTypeRepository
	PublishedPostRepo repository.PublishedPostRepository
}

func (app *PostTypeApp) GetPostTypeDetail(slug string) (*entity.PostType, error) {
	pt, err := app.PostTypeRepo.GetPostTypeBySlug(slug)
	if err != nil {
		return nil, err
	}

	postCount, err := app.PublishedPostRepo.GetPublishedPostCountByPostTypeId(pt.ID)
	if err != nil {
		return nil, err
	}

	pt.PublishedPostCount = postCount

	return pt, nil
}
