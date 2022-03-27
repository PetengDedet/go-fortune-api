package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type MediaRepository interface {
	GetMediaByIds(ids []int64) ([]entity.Media, error)
}
