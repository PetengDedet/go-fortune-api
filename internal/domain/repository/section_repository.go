package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type SectionRepository interface {
	GetSectionsByPageId(pageId int64) ([]entity.Section, error)
}
