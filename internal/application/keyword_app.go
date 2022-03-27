package application

import (
	"errors"

	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type KeywordAppInterface interface {
	SaveKeyword(keyword string) (bool, error)
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
