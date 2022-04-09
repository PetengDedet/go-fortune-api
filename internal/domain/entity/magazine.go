package entity

import "gopkg.in/guregu/null.v4"

type Magazine struct {
	ID           int64          `json:"-"`
	Title        string         `json:"title"`
	Slug         string         `json:"slug"`
	Edition      null.String    `json:"edition"`
	Description  null.String    `json:"description"`
	PurchaseLink null.String    `json:"purchase_link"`
	Cover        *MagazineCover `json:"cover"`

	Media *Media `json:"-"`
}

type MagazineCover struct {
	ImageUrl    null.String `json:"image_url"`
	SourceName  null.String `json:"source_name"`
	SourceUrl   null.String `json:"source_url"`
	Description null.String `json:"description"`
	EmbedVideo  null.String `json:"embed_video"`
}
