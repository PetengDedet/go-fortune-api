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

	MenuPositionID  int64       `json:"-"`
	TableID         null.Int    `json:"-"`
	TableName       null.String `json:"-"`
	GeneralStatusID int64       `json:"-"`

	Category *Category `json:"-"`
	Page     *Page     `json:"-"`
	Linkout  *Linkout  `json:"-"`
	Rank     *Rank     `json:"-"`
}
