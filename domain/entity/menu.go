package entity

import "gopkg.in/guregu/null.v4"

type MenuPayload struct {
	ID              int64       `db:"db"`
	OrderNum        int64       `db:"order_num"`
	Title           string      `db:"title"`
	Slug            string      `db:"slug"`
	ParentMenuID    null.Int    `db:"parent_menu_id"`
	GeneralStatusID null.Int    `db:"general_status_id"`
	MenuPositionID  null.Int    `db:"menu_position_id"`
	TableName       null.Int    `db:"table_name"`
	TableID         null.Int    `db:"table_id"`
	CreatedAt       null.Time   `db:"created_at"`
	CreatedBy       null.Int    `db:"created_by"`
	CategoryName    null.String `db:"category_name"`
	CategorySlug    null.String `db:"category_slug"`
	PageName        null.String `db:"page_name"`
	PageSlug        null.String `db:"page_slug"`
	LinkoutType     null.String `db:"linkout_type"`
	LinkoutUrl      null.String `db:"linkout_url"`
	RankTitle       null.String `db:"rank_title"`
	RankName        null.String `db:"rank_name"`
	RankSlug        null.String `db:"rank_slug"`
}

type Menu struct {
	ID        int64       `json:"-"`
	Name      null.String `json:"name"`
	Slug      null.String `json:"slug"`
	Type      null.String `json:"type"`
	Url       null.String `json:"url"`
	OrdeNum   int64       `json:"order_num"`
	IsActive  bool        `json:"is_active"`
	ChildMenu []Menu      `json:"child_menu"`

	Category *Category `json:"-"`
	Page     *Page     `json:"-"`
	Linkout  *Linkout  `json:"-"`
	Rank     *Rank     `json:"-"`

	Payload *MenuPayload `json:"-"`
}

func (p *MenuPayload) MenuResponse() *Menu {
	return nil
}

type CategoryMenu struct {
	Menu    *Menu
	Payload *MenuPayload
}
