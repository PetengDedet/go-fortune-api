package entity

import "gopkg.in/guregu/null.v4"

type Page struct {
	Slug            null.String `json:"slug" db:"slug"`
	Excerpt         null.String `json:"excerpt" db:"excerpt"`
	Description     null.String `json:"description" db:"description"`
	Url             null.String `json:"url" db:"url"`
	MetaTitle       null.String `json:"meta_title" db:"meta_title"`
	MetaDescription null.String `json:"meta_description" db:"meta_description"`

	Page          null.String `json:"page" db:"-"`
	ArticleCounts int64       `json:"article_counts" db:"-"`
	Sections      []Section   `json:"sections" db:"-"`

	ID              int64       `json:"-" db:"id"`
	MediaID         null.Int    `json:"-" db:"media_id"`
	Name            null.String `json:"-" db:"name"`
	GeneralStatusID int64       `json:"-" db:"general_status_id"`
	CreatedAt       null.Time   `json:"-" db:"created_at"`
	UpdatedAt       null.Time   `json:"-" db:"updated_at"`
	CreatedBy       int64       `json:"-" db:"created_by"`
	UpdatedBy       null.Int    `json:"-" db:"updated_by"`

	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`

	Articles *SearchResultArticles `json:"articles,omitempty" db:"-"`
}
