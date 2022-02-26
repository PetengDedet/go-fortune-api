package entity

import (
	"os"
	"strconv"
)

type MenuPosition struct {
	ID   int64
	Name string
	Slug string
}

type PublicMenuPosition struct {
	Position string       `json:"position"`
	Menus    []PublicMenu `json:"menus"`
}

func PublicMenuPositionResponse(mp MenuPosition, parentMenus []Menu, childrenMenus []Menu) *PublicMenuPosition {
	var pm []PublicMenu
	for _, parMen := range parentMenus {
		if parMen.MenuPositionID == mp.ID {
			var cm []PublicMenu
			for _, childMen := range childrenMenus {
				if childMen.ParentMenuID == parMen.ID {
					isActive := false
					if childMen.IsActive == 1 {
						isActive = true
					}

					cm = append(cm, PublicMenuResponse(
						PublicMenu{
							Name:       childMen.Title,
							Slug:       childMen.Slug,
							Type:       childMen.MenuType,
							Url:        childMen.Slug,
							OrderNum:   childMen.OrderNum,
							LinkoutUrl: childMen.LinkoutUrl,
							IsActive:   isActive,
							ChildMenu:  []PublicMenu{},
						}))
				}
			}
			isActive := false
			if parMen.IsActive == 1 {
				isActive = true
			}
			pm = append(pm, PublicMenuResponse(
				PublicMenu{
					Name:       parMen.Title,
					Slug:       parMen.Slug,
					Type:       parMen.MenuType,
					Url:        parMen.Slug,
					OrderNum:   parMen.OrderNum,
					LinkoutUrl: parMen.LinkoutUrl,
					IsActive:   isActive,
					ChildMenu:  cm,
				}))
		}
	}

	if mp.Slug == "header" {
		headerLimit, err := strconv.Atoi(os.Getenv("HEADER_MENU_LIMIT"))
		if err != nil {
			headerLimit = 9
		}

		pm = pm[:headerLimit]
	}

	return &PublicMenuPosition{
		Position: mp.Slug,
		Menus:    pm,
	}
}
