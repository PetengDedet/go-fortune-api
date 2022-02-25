package entity

type MenuPosition struct {
	ID   int64
	Name string
	Slug string
}

type PublicMenuPosition struct {
	Position string
	Menus    []PublicMenu
}

// func PublicMenuPositionResponse(mp *MenuPosition, parentMenus []ParentMenu, childrenMenus []ChildrenMenu) *PublicMenuPosition {
// 	var pm []PublicMenu
// 	for i, parMen := range parentMenus {
// 		if parMen.MenuPositionID == mp.ID {
// 			pm[i] = PublicMenu{

// 			}
// 		}
// 	}
// }
