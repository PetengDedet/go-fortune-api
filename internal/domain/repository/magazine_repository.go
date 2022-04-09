package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type MagazineRepository interface {
	GetLatestActiveMagazines(limit int) ([]entity.Magazine, error)
}
