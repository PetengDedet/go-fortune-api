package mysql

import (
	"fmt"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
	"github.com/jmoiron/sqlx"
)

type MenuRepo struct {
	db *sqlx.DB
}

func NewMenuRepository(db *sqlx.DB) *MenuRepo {
	return &MenuRepo{db}
}

//MenuRepo implements the repository.MenuRepository interface
var _ repository.MenuRepository = &MenuRepo{}

func (r *MenuRepo) GetMenuPositions() ([]entity.MenuPosition, error) {

	menuPositions := []entity.MenuPosition{}
	rows, err := r.db.Queryx("SELECT id, name, slug FROM menu_positions")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var mp entity.MenuPosition
		err = rows.Scan(&mp.ID, &mp.Name, &mp.Slug)
		fmt.Println(err)
		fmt.Println(mp)
	}

	fmt.Println(menuPositions)

	return nil, rows.Err()
}

func (r *MenuRepo) GetMenusByPositionIds(positionIds []int) ([]entity.Menu, error)
func (r *MenuRepo) GetParentMenus(menuIds []int) ([]entity.Menu, error)
