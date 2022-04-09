package application

import (
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type MagazineApp struct {
	MagazineRepo repository.MagazineRepository
}

func (app *MagazineApp) GetLatestMagazines() ([]entity.Magazine, error) {
	limit, err := strconv.Atoi(os.Getenv("HOMEPAGE_MAGAZINE_LIMIT"))
	if err != nil {
		limit = 10
	}

	magazines, err := app.MagazineRepo.GetLatestActiveMagazines(limit)
	if err != nil {
		return nil, err
	}

	return magazines, nil
}
