package application

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type MenuApp struct {
	MenuRepo repository.MenuRepository
}

func (ma *MenuApp) GetPublicMenuPositions() ([]entity.PublicMenuPosition, error) {
	menuPost, err := ma.MenuRepo.GetMenuPositions()
	if err != nil {
		panic(err.Error())
	}

	var positionIds []int
	for _, mp := range menuPost {
		positionIds = append(positionIds, int(mp.ID))
	}

	log.Println("positionIds", positionIds)
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

	log.Println("menuPos", menuPost)
	log.Println("parentMenus:", parentMenus)
	log.Println("childrenMenus", childrenMenus)

	return nil, nil
}
