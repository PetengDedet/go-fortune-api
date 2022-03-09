package mysql

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type MenuRepo struct {
	DB *sqlx.DB
}

func (menuRepo *MenuRepo) GetMenuPositions() ([]entity.MenuPosition, error) {
	query := `
		SELECT 
			mp.id,
			mp.name,
			mp.slug
		FROM menu_positions mp
	`

	rows, err := menuRepo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var menuPositions []entity.MenuPosition
	for rows.Next() {
		menuPosition := &entity.MenuPosition{}
		var id int64
		var name string
		var slug string

		err := rows.Scan(&id, &name, &slug)
		if err != nil {
			return nil, err
		}

		menuPosition.ID = id
		menuPosition.Name = name
		menuPosition.Slug = slug

		menuPositions = append(menuPositions, *menuPosition)
	}

	return menuPositions, nil
}

func (menuRepo *MenuRepo) GetMenusByPositionIds(positionIds []int) ([]entity.Menu, error) {
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
			
			c.id AS category_id,
			c.name AS category_name,
			c.slug AS category_slug,
			
			p.id AS page_id,
			p.name AS page_name,
			p.slug AS page_slug,

			lo.id AS linkout_id,
			lo.type AS linkout_type,
			lo.url AS linkout_url

		FROM menus m
		LEFT JOIN categories c ON c.id = m.table_id AND m.table_id = 'categories'
		LEFT JOIN pages p ON p.id = m.table_id AND m.table_id = 'pages'
		LEFT JOIN linkouts lo ON lo.id = m.table_id AND m.table_name = 'linkouts'
		WHERE ISNULL(m.parent_menu_id)
			AND m.menu_position_id IN(?)
			AND m.general_status_id = ?
	`, positionIds, 1)

	if err != nil {
		return nil, err
	}

	query = menuRepo.DB.Rebind(query)
	rows, err := menuRepo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var menus []entity.Menu
	for rows.Next() {
		var category entity.Category
		var page entity.Page
		var linkout entity.Linkout
		var menu entity.Menu
		err := rows.Scan(
			&menu.ID,
			&menu.Name,
			&menu.Slug,
			&menu.OrdeNum,
			&menu.MenuPositionID,
			&menu.TableID,
			&menu.TableName,
			&menu.GeneralStatusID,

			&category.ID,
			&category.Name,
			&category.Slug,

			&page.ID,
			&page.Name,
			&page.Slug,

			&linkout.ID,
			&linkout.Type,
			&linkout.Url,
		)
		if err != nil {
			return nil, err
		}

	}

	return menus, nil
}

func (menuRepo *MenuRepo) GetChildrenMenus(menuIds []int) ([]entity.Menu, error) {
	return nil, nil
	// query, args, err := sqlx.In(`
	// 	SELECT
	// 		m.id,
	// 		m.title,
	// 		m.slug,
	// 		m.parent_menu_id,
	// 		m.order_num,
	// 		m.menu_position_id,
	// 		m.table_name AS menu_type,
	// 		m.general_status_id AS is_active,
	// 		COALESCE(lo.url, '') AS linkout_url
	// 	FROM menus m
	// 	LEFT JOIN linkouts lo ON lo.id = m.table_id AND m.table_name = 'linkouts'
	// 	WHERE m.parent_menu_id IN(?)
	// 		AND m.general_status_id = ?
	// `, menuIds, 1)

	// if err != nil {
	// 	return nil, err
	// }

	// query = menuRepo.DB.Rebind(query)
	// rows, err := menuRepo.DB.Query(query, args...)
	// if err != nil {
	// 	return nil, err
	// }

	// var menus []entity.Menu
	// for rows.Next() {
	// 	var menu entity.Menu
	// 	var id int64
	// 	var name string
	// 	var slug string
	// 	var parent_menu_id int64
	// 	var order_num int64
	// 	var menu_position_id int64
	// 	var menu_type string
	// 	var is_active int64
	// 	var linkout_url string

	// 	err := rows.Scan(
	// 		&id,
	// 		&name,
	// 		&slug,
	// 		&parent_menu_id,
	// 		&order_num,
	// 		&menu_position_id,
	// 		&menu_type,
	// 		&is_active,
	// 		&linkout_url,
	// 	)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	menu.ID = id
	// 	menu.Title = name
	// 	menu.Slug = slug
	// 	menu.ParentMenuID = parent_menu_id
	// 	menu.OrderNum = order_num
	// 	menu.MenuPositionID = menu_position_id
	// 	menu.MenuType = menu_type
	// 	menu.IsActive = is_active
	// 	menu.LinkoutUrl = linkout_url

	// 	menus = append(menus, menu)
	// }

	// return menus, nil
}
