package application

import (
	"errors"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type KeywordAppInterface interface {
	SaveKeyword(keyword string) (bool, error)
	GetPopularKeywords() ([]entity.Keyword, error)
}

type KeywordApp struct {
	KeywordRepo repository.KeywordRepository
}

func (app *KeywordApp) SaveKeyword(keyword string) error {
	if len(keyword) == 0 {
		return errors.New("keyword can not be empty")
	}

	return app.KeywordRepo.SaveNewKeyword(keyword)
}

func (app *KeywordApp) GetPopularKeyword() ([]entity.Keyword, error) {
	return app.KeywordRepo.GetPopularKeyword()
}
