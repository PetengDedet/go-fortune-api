package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type MenuRepository interface {
	GetMenuPositions() ([]entity.MenuPosition, error)
	GetMenusByPositionIds(positionIds []int) ([]entity.Menu, error)
	GetChildrenMenus(menuIds []int) ([]entity.Menu, error)
}
