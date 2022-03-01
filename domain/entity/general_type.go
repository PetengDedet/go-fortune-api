package entity

type GeneralType struct {
	ID   int64  `json:"id" db:"name"`
	Name string `json:"name" db:"name"`
	Slug string `json:"slug" db:"slug"`
}
