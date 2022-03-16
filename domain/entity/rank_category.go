package entity

import "gopkg.in/guregu/null.v4"

type RankCategory struct {
	ID              int64       `json:"-"`
	Name            string      `json:"name"`
	Slug            string      `json:"slug"`
	Excerpt         null.String `json:"excerpt"`
	MetaTitle       null.String `json:"meta_title"`
	MetaDescription null.String `json:"meta_description"`
	IsActive        bool        `json:"is_active"`

	CreatedAt null.Time `json:"-" db:"created_at"`
	UpdatedAt null.Time `json:"-" db:"updated_at"`
	CreatedBy int64     `json:"-" db:"created_by"`
	UpdatedBy null.Int  `json:"-" db:"updated_by"`

	GeneralType   *GeneralType   `json:"-" db:"-"`
	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`
	Media         *Media         `json:"-" db:"-"`
}
