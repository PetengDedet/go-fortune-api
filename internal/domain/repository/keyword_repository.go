package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type KeywordRepository interface {
	SaveNewKeyword(keyword string) error
	GetPopularKeyword() ([]entity.Keyword, error)
}
