package entity

import "gopkg.in/guregu/null.v4"

type Category struct {
	Name            string      `json:"name" db:"name"`
	Slug            string      `json:"slug" db:"slug"`
	Excerpt         null.String `json:"excerpt" db:"excerpt"`
	MetaTitle       null.String `json:"meta_title" db:"meta_title"`
	MetaDescription null.String `json:"meta_description" db:"meta_description"`

	ID              int64    `json:"-" db:"id"`
	GeneralStatusID int64    `json:"-" db:"general_status_id"`
	MediaID         null.Int `json:"-" db:"media_id"`

	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Media         *Media         `json:"media" db:"-"`
}
