package entity

import "gopkg.in/guregu/null.v4"

type PageSection struct {
	ID       int64 `json:"-"`
	Type     null.String
	Title    null.String
	OrderNum int64
	Url      null.String
	BaseUrl  null.String
	ImageUrl null.String
}
