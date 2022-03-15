package entity

import "gopkg.in/guregu/null.v4"

type Media struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name" db:"name"`
	Description null.String `json:"descrtiption" db:"description"`
	UrlMedia    null.String `json:"url_media" db:"url_media"`
	Mime        null.String `json:"mime" db:"mime"`
	Extension   null.String `json:"extension" db:"extension"`
	SourceName  null.String `json:"source_name" db:"source_name"`
	SourceUrl   null.String `json:"source_url" db:"source_url"`
	Width       int64       `json:"width" db:"width"`
	Height      int64       `json:"height" db:"height"`
	Keyword     null.String `json:"keyword" db:"keyword"`
	UrlEmbed    null.String `json:"url_embed" db:"url_embed"`

	GeneralStatusID int64     `json:"-" db:"general_status_id"`
	GeneralTypeID   int64     `json:"-" db:"general_type_id"`
	GalleryID       null.Int  `json:"-" db:"gallery_id"`
	CreatedAt       null.Time `json:"-" db:"created_at"`
	UpdatedAt       null.Time `json:"-" db:"updated_at"`
	CreatedBy       int64     `json:"-" db:"created_by"`
	UpdatedBy       null.Int  `json:"-" db:"updated_by"`

	GeneralType   *GeneralType   `json:"-" db:"-"`
	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Gallery       *Gallery       `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`
}
