package entity

import "gopkg.in/guregu/null.v4"

type Category struct {
	ID              int64
	GeneralStatusID int64
	Name            string
	Slug            string
	Excerpt         string
	MediaID         int64
	MetaTitle       string
	MetaDescription string
}

type CategoryPage struct {
	ID              int64       `json:"-"`
	Page            null.String `json:"page"`
	Slug            null.String `json:"slug"`
	Excerpt         null.String `json:"excerpt"`
	ArticleCounts   null.Int    `json:"article_counts"`
	Description     null.String `json:"description"`
	Url             null.String `json:"url"`
	MetaTitle       null.String `json:"meta_title"`
	MetaDescription null.String `json:"meta_description"`
	Sections        []Section   `json:"sections"`
}
