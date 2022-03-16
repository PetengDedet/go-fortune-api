package repository

import "github.com/PetengDedet/fortune-post-api/domain/entity"

type TagRepository interface {
	GetTagByIds(ids []int64) ([]entity.Tag, error)
}
