package entity

import (
	"os"
	"strconv"

	"gopkg.in/guregu/null.v4"
)

type MenuPosition struct {
	ID       int64  `json:"-"`
	Name     string `json:"-"`
	Position string `json:"position"`
	Slug     string `json:"-"`
	Menus    []Menu `json:"menus"`
}

func (menuPosition *MenuPosition) GetMenus(menus []Menu) *MenuPosition {
	var mns []Menu
	for _, m := range menus {
		if m.MenuPositionID == menuPosition.ID {
			mns = append(mns, m)
		}
	}

	if menuPosition.Slug == "header" {
		headerLimit, err := strconv.Atoi(os.Getenv("HEADER_MENU_LIMIT"))
		if err != nil {
			headerLimit = 9
		}

		if len(mns) > headerLimit {
			tmpMenus := mns
			moreChildMenus := tmpMenus[headerLimit:]
			moreMenu := Menu{
				Name:      null.StringFrom("MORE"),
				Slug:      null.StringFrom("more"),
				OrderNum:  int64(headerLimit) + 1,
				IsActive:  true,
				ChildMenu: moreChildMenus,
			}

			mns = mns[:headerLimit]
			mns = append(mns, moreMenu)
		}
	}
	menuPosition.Menus = mns

	return menuPosition
}
