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

// func (pmp *PublicMenuPositions) PublicMenuPositions(mp *MenuPosition, m *Menus, p *Menus) []interface{} {
// 	result := make([]interface{}, len(mp))
// 	for index, menuPosition := range mp {
// 		result[index] = user.PublicUser()
// 	}
// 	return result
// }
