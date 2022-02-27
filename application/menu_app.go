package application

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type MenuApp struct {
	MenuRepo repository.MenuRepository
}

func (ma *MenuApp) GetPublicMenuPositions() ([]entity.PublicMenuPosition, error) {
	var menuPositions []entity.PublicMenuPosition

	menuPost, err := ma.MenuRepo.GetMenuPositions()
	if err != nil {
		return nil, err
	}

	if len(menuPost) == 0 {
		return []entity.PublicMenuPosition{}, nil
	}

	var positionIds []int
	for _, mp := range menuPost {
		positionIds = append(positionIds, int(mp.ID))
	}

	parentMenus, err := ma.MenuRepo.GetMenusByPositionIds(positionIds)
	if err != nil {
		return nil, err
	}

	// No Parent menus, just return the menu positions
	if len(parentMenus) == 0 {
		return entity.PublicMenuPositionsResponse(menuPost, nil, nil), nil
	}

	var parentMenuIds []int
	for _, mp := range parentMenus {
		parentMenuIds = append(parentMenuIds, int(mp.ID))
	}

	childrenMenus, err := ma.MenuRepo.GetChildrenMenus(parentMenuIds)
	if err != nil {
		return nil, err
	}

	menuPositions = entity.PublicMenuPositionsResponse(menuPost, parentMenus, childrenMenus)

	return menuPositions, nil
}
