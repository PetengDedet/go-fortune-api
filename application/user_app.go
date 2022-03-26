package application

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type UserAppInterface interface {
	GetAuthorsOfPostByPostIds(postIds []int64) ([]entity.Author, error)
}

type UserApp struct {
	UserRepo repository.UserRepository
}

func (app *UserApp) GetAuthorsOfPostByPostIds(postIds []int64) ([]entity.Author, error) {
	return app.UserRepo.GetAuthorsByPostIds(postIds)
}
