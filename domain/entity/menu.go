package entity

import "gopkg.in/guregu/null.v4"

type Menu struct {
	ID        int64       `json:"-"`
	Name      null.String `json:"name"`
	Slug      null.String `json:"slug"`
	Type      null.String `json:"type"`
	Url       null.String `json:"url"`
	OrdeNum   int64       `json:"order_num"`
	IsActive  bool        `json:"is_active"`
	ChildMenu []Menu      `json:"child_menu"`

	Title           null.String `json:"-"`
	MenuPositionID  int64       `json:"-"`
	TableID         null.Int    `json:"-"`
	TableName       null.String `json:"-"`
	GeneralStatusID int64       `json:"-"`
	ParentMenuId    null.Int    `json:"-"`

	Category *Category `json:"-"`
	Page     *Page     `json:"-"`
	Linkout  *Linkout  `json:"-"`
	Rank     *Rank     `json:"-"`
}

func (m *Menu) ClassifyMenu() *Menu {
	if m.Category != nil {
		m.Type = null.StringFrom("category")
		m.Url = null.StringFrom("/" + m.Slug.String)

		return m
	}

	if m.Linkout != nil {
		m.Type = null.StringFrom("link")
		m.Url = null.StringFrom(m.Linkout.Url)

		return m
	}

	if m.Page != nil {
		m.Type = null.StringFrom("page")
		m.Url = m.Page.Url

		return m
	}

	if m.Rank != nil {
		m.Type = null.StringFrom("rank")
		m.Url = null.StringFrom("/" + m.Rank.Slug)

		return m
	}

	return m
}

func (m *Menu) GetChildMenus(childrens []Menu) []Menu {
	var childs []Menu
	for _, c := range childrens {
		if c.ParentMenuId.Int64 == m.ID {
			childs = append(childs, c)
		}
	}

	return childs
}
