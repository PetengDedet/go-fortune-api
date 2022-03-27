package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type LinkoutRepository interface {
	GetLinkoutsByIds(linkoutIds []int64) ([]entity.Linkout, error)
}
