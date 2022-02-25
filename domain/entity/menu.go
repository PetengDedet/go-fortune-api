package entity

type PublicMenu struct {
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	Type      string       `json:"type"`
	Url       string       `json:"url"`
	OrderNum  int64        `json:"order_num"`
	IsActive  int64        `json:"is_active"`
	ChildMenu []PublicMenu `json:"child_menu"`
}

type ParentMenu struct {
	ID             int64
	Title          string
	Slug           string
	OrderNum       int64
	MenuPositionID int64
	MenuType       string
}

type ChildrenMenu struct {
	ID             int64
	Title          string
	Slug           string
	ParentMenuID   int64
	OrderNum       int64
	MenuPositionID int64
	MenuType       string
}
