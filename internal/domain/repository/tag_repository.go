package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type TagRepository interface {
	GetTagByIds(ids []int64) ([]entity.Tag, error)
	GetTagBySlug(slug string) (*entity.Tag, error)
}
