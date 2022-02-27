package repository

import "github.com/PetengDedet/fortune-post-api/domain/entity"

type PageRepository interface {
	GetPageBySlug(slug string) (*entity.Page, error)
	GetSectionsByPageId(pageId int64) ([]entity.Section, error)
}
