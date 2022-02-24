package entity

type PublicMenu struct {
	Name      string       `json:"name"`
	Slug      string       `json:"slug"`
	Type      string       `json:"type"`
	Url       string       `json:"url"`
	OrderNum  int          `json:"order_num"`
	IsActive  int          `json:"is_active"`
	ChildMenu []PublicMenu `json:"child_menu"`
}

type Menu struct {
	ID             int
	Title          string
	Slug           string
	ParentMenuID   int
	OrderNum       int
	MenuPositionID int
	MenuType       string
}
