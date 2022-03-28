package entity

import "time"

type Keyword string
type KeywordHistory struct {
	Keyword   string     `json:"keyword" bson:"keyword"`
	CreatedAt *time.Time `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}
