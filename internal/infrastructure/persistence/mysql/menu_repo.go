package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type MenuRepo struct {
	DB *sqlx.DB
}

func (repo *MenuRepo) GetMenuPositions() ([]entity.MenuPosition, error) {
	query := `
		SELECT 
			mp.id,
			mp.name,
			mp.slug
		FROM menu_positions mp
	`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var menuPositions []entity.MenuPosition
	for rows.Next() {
		var menuPosition entity.MenuPosition
		err := rows.Scan(
			&menuPosition.ID,
			&menuPosition.Position,
			&menuPosition.Slug,
		)
		if err != nil {
			return nil, err
		}

		menuPositions = append(menuPositions, menuPosition)
	}

	return menuPositions, nil
}

func (repo *MenuRepo) GetMenusByPositionIds(positionIds []int) ([]entity.Menu, error) {
	query, args, err := sqlx.In(`
		SELECT 
			m.id,
			m.title,
			m.slug,
			m.order_num,
			m.menu_position_id,
			m.table_id,
			m.table_name,
			m.general_status_id
			
		FROM menus m
		WHERE ISNULL(m.parent_menu_id)
			AND m.menu_position_id IN(?)
			AND m.general_status_id = ?
		ORDER BY order_num ASC
	`, positionIds, 1)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var menus []entity.Menu
	for rows.Next() {
		var menu entity.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Slug,
			&menu.OrderNum,
			&menu.MenuPositionID,
			&menu.TableID,
			&menu.TableName,
			&menu.GeneralStatusID,
		)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}

func (repo *MenuRepo) GetChildrenMenus(menuIds []int) ([]entity.Menu, error) {
	query, args, err := sqlx.In(`
		SELECT 
			m.id,
			m.title,
			m.slug,
			m.order_num,
			m.menu_position_id,
			m.table_id,
			m.table_name,
			m.general_status_id,
			m.parent_menu_id
			
		FROM menus m
		WHERE m.parent_menu_id IN(?)
			AND m.general_status_id = ?
		ORDER BY m.order_num ASC
	`, menuIds, 1)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var menus []entity.Menu
	for rows.Next() {
		var menu entity.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Slug,
			&menu.OrderNum,
			&menu.MenuPositionID,
			&menu.TableID,
			&menu.TableName,
			&menu.GeneralStatusID,
			&menu.ParentMenuId,
		)
		if err != nil {
			return nil, err
		}

		menus = append(menus, menu)
	}

	return menus, nil
}
