package repository

import "github.com/PetengDedet/fortune-post-api/domain/entity"

type PostTypeRepository interface {
	GetPostTypeByIds(ids []int64) ([]entity.PostType, error)
	GetPostTypeBySlug(slug string) (*entity.PostType, error)
}
