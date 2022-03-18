package entity

import "gopkg.in/guregu/null.v4"

type Tag struct {
	ID              int64       `json:"-" db:"id"`
	Name            string      `json:"name" db:"name"`
	Slug            string      `json:"slug" db:"slug"`
	Excerpt         null.String `json:"excerpt" db:"excerpt"`
	MetaTitle       null.String `json:"meta_title" db:"meta_title"`
	MetaDescription null.String `json:"meta_description" db:"meta_description"`

	MediaID null.Int `json:"-" db:"media_id"`

	Media              *Media `json:"-" db:"-"`
	PublishedPostCount int64  `json:"-" db:"-"`
}
