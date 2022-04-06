package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type UserRepository interface {
	GetAuthorsByPostIds(postIds []int64) ([]entity.Author, error)
	GetAuthorsByPostId(postId int64) ([]entity.Author, error)
}
