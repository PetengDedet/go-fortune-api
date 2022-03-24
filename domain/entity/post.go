package entity

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Post struct {
	ID              int64       `json:"id" db:"id"`
	CategoryID      int64       `json:"category_id" db:"category_id"`
	Title           string      `json:"title" db:"title"`
	Slug            string      `json:"slug" db:"slug"`
	Description     string      `json:"description" db:"description"`
	Excerpt         string      `json:"excerpt" db:"excerpt"`
	IsCSC           bool        `json:"is_csc" db:"is_csc"`
	VisitedCount    int64       `json:"visited_count" db:"visited_count"`
	PublishAt       null.Time   `json:"publish_at" db:"publish_at"`
	MetaTitle       null.String `json:"meta_title" db:"meta_title"`
	MetaDescription null.String `json:"meta_description" db:"meta_description"`
	OGCaption       null.String `json:"og_caption" db:"og_caption"`
	OGTitle         null.String `json:"og_title" db:"og_title"`
	OGDescription   null.String `json:"og_description" db:"og_description"`
	UUID            null.String `json:"uuid" db:"uuid"`

	PostTypeID   string      `json:"-" db:"post_type_id"`
	PostStatusID string      `json:"-" db:"post_status_id"`
	CampaignID   null.String `json:"-" db:"campaign_id"`

	CreatedAt    *time.Time `json:"-" db:"created_at"`
	UpdatedAt    null.Time  `json:"-" db:"updated_at"`
	CreatedBy    int64      `json:"-" db:"created_by"`
	UpdatedBy    null.Int   `json:"-" db:"updated_by"`
	CoverMediaID int64      `json:"-" db:"cover_media_id"`
	LatestBatch  int64      `json:"-" db:"latest_batch"`

	Creator *User `json:"-" db:"-"`
	Updater *User `json:"-" db:"-"`

	Category *Category `json:"category" db:"-"`
	Tags     []Tag     `json:"tags" db:"-"`
	Cover    *Media    `json:"cover" db:"-"`
}
