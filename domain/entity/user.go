package entity

import "gopkg.in/guregu/null.v4"

type User struct {
	ID          int64       `json:"id" db:"id"`
	UUID        null.String `json:"uuid" db:"uuid"`
	Email       null.String `json:"email" db:"email"`
	Name        string      `json:"name" db:"name"`
	Username    string      `json:"username" db:"username"`
	Nickname    null.String `json:"nickname" db:"nickname"`
	BirthDate   null.Time   `json:"birth_date" db:"birth_date"`
	Description null.String `json:"description" db:"description"`
	Avatar      null.String `json:"avatar" db:"avatar"`
	Banner      null.Int    `json:"banner" db:"banner"`
	Gender      null.Int    `json:"gender" db:"gender"`

	Password        null.String `json:"-" db:"password"`
	CreatedAt       null.Time   `json:"-" db:"created_at"`
	UpdatedAt       null.Time   `json:"-" db:"updated_at"`
	CreatedBy       null.Int    `json:"-" db:"created_by"`
	UpdatedBy       null.Int    `json:"-" db:"updated_by"`
	GeneralStatusID int64       `json:"-" db:"general_status_id"`
	UserTypeID      int64       `json:"-" db:"user_type_id"`

	UserType      *UserType      `json:"-" db:"-"`
	GeneralStatus *GeneralStatus `json:"-" db:"-"`
}
