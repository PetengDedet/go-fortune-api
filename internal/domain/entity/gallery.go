package entity

import "gopkg.in/guregu/null.v4"

type Gallery struct {
	ID int64 `json:"id" db:"id"`

	Name        string      `json:"name" db:"name"`
	Description null.String `json:"description" db:"description"`
	Path        null.String `json:"path" db:"path"`
	IsPublic    null.Int    `json:"is_public" db:"is_public"`

	GeneralStatusID int64     `json:"-" db:"general_status_id"`
	CreatedAt       null.Time `json:"-" db:"created_at"`
	UpdatedAt       null.Time `json:"-" db:"updated_at"`
	CreatedBy       int64     `json:"-" db:"created_by"`
	UpdatedBy       null.Int  `json:"-" db:"updated_by"`

	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`
}
