package application

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type menuApp struct {
	menuRepo repository.MenuRepository
}

//MenuApp implements the MenuAppInterface
var _ MenuAppInterface = &menuApp{}

type MenuAppInterface interface {
	GetMenuPositions() ([]entity.PublicMenuPosition, error)
}

func (app *menuApp) GetMenuPositions() ([]entity.PublicMenuPosition, error) {
	menuPositions, err := app.menuRepo.GetMenuPositions()
	if err != nil {
		return nil, err
	}

	var menuPositionIds []int
	for index, m := range menuPositions {
		menuPositionIds[index] = int(m.ID)
	}

	menus, err := app.menuRepo.GetMenusByPositionIds(menuPositionIds)
	if err != nil {
		return nil, err
	}

	var menuIds []int
	for index, m := range menus {
		menuIds[index] = m.ID
	}

	parentMenu, err := app.menuRepo.GetParentMenus(menuIds)
	if err != nil {
		return nil, err
	}

	var menuPos []entity.PublicMenuPosition
	// for index, mp := range menuPositions {
	// 	// var parent *enti
	// 	// for i, parent := range
	// 	// menuPos[index] = entity.PublicMenuPosition{
	// 	// 	Position: mp.Name,
	// 	// }
	// }

	return menuPos, nil
}
