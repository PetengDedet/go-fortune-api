package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type PageRepository interface {
	GetPageBySlug(slug string) (*entity.Page, error)
	GetPagesByIds(pageIds []int64) ([]entity.Page, error)
}
