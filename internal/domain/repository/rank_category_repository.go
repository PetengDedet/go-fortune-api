package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type RankCategoryRepository interface {
	GetRankCategoryByIds(ids []int64) ([]entity.RankCategory, error)
}
