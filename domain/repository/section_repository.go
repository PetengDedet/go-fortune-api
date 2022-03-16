package repository

import "github.com/PetengDedet/fortune-post-api/domain/entity"

type SectionRepository interface {
	GetSectionsByPageId(pageId int64) ([]entity.Section, error)
	// GetSectionsByPageSlug(slug string) ([]entity.Section, error)
}
