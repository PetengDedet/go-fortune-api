package entity

import "gopkg.in/guregu/null.v4"

type Linkout struct {
	ID   int64  `json:"id" db:"id"`
	Url  string `json:"url" db:"url"`
	Type string `json:"type" db:"type"`

	CreatedAt null.Time `json:"-" db:"created_at"`
	CreatedBy int64     `json:"-" db:"created_by"`

	Creator *User `json:"-" db:"-"`
}
