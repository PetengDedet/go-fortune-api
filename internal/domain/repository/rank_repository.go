package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type RankRepository interface {
	GetRanksByIds(ids []int64) ([]entity.Rank, error)
}
