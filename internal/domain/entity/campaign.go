package entity

import "gopkg.in/guregu/null.v4"

type Campaign struct {
	ID   int64  `json:"-" db:"id"`
	Name string `json:"name" db:"name"`
	Slug string `json:"slug" db:"slug"`

	MediaID         null.Int  `json:"-" db:"media_id"`
	GeneralStatusID int64     `json:"-" db:"general_status_id"`
	CreatedAt       null.Time `json:"-" db:"created_at"`
	UpdatedAt       null.Time `json:"-" db:"updated_at"`
	CreatedBy       int64     `json:"-" db:"created_by"`
	UpdatedBy       null.Int  `json:"-" db:"updated_by"`

	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`
}
