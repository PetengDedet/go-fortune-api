package application

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type MenuApp struct {
	MenuRepo repository.MenuRepository
}

func (ma *MenuApp) GetPublicMenuPositions() []entity.PublicMenuPosition {
	menuPost, err := ma.MenuRepo.GetMenuPositions()
	if err != nil {
		panic(err.Error())
	}

	var positionIds []int
	for _, mp := range menuPost {
		positionIds = append(positionIds, int(mp.ID))
	}

	parentMenus, err := ma.MenuRepo.GetMenusByPositionIds(positionIds)
	if err != nil {
		panic(err.Error())
	}

	var parentMenuIds []int
	for _, mp := range parentMenus {
		parentMenuIds = append(parentMenuIds, int(mp.ID))
	}

	childrenMenus, err := ma.MenuRepo.GetChildrenMenus(parentMenuIds)
	if err != nil {
		panic(err.Error())
	}

	var publicMenuPositions []entity.PublicMenuPosition
	for _, pmp := range menuPost {
		publicMenuPositions = append(publicMenuPositions, *entity.PublicMenuPositionResponse(pmp, parentMenus, childrenMenus))
	}

	return publicMenuPositions
}
