package entity

import "gopkg.in/guregu/null.v4"

type PostType struct {
	ID              int64       `json:"-" db:"id"`
	Name            string      `json:"name" db:"name"`
	Slug            string      `json:"slug" db:"slug"`
	Excerpt         null.String `json:"excerpt,omitempty" db:"excerpt"`
	MetaTitle       null.String `json:"meta_title,omitempty" db:"meta_title"`
	MetaDescription null.String `json:"meta_description,omitempty" db:"meta_description"`

	MediaID null.Int `json:"-" db:"media_id"`

	Media              *Media `json:"-" db:"-"`
	PublishedPostCount int64  `json:"-" db:"-"`
}

type SearchResultPostType struct {
	Name string `json:"name" db:"name"`
	Slug string `json:"slug" db:"slug"`
}
