package entity

import "gopkg.in/guregu/null.v4"

type Page struct {
	ID              int64         `json:"-"`
	Page            null.String   `json:"page"`
	Slug            null.String   `json:"slug"`
	Excerpt         null.String   `json:"excerpt"`
	ArticleCounts   int64         `json:"article_counts"`
	Description     null.String   `json:"description"`
	Url             null.String   `json:"url"`
	MetaTitle       null.String   `json:"meta_title"`
	MetaDescription null.String   `json:"meta_description"`
	Sections        []PageSection `json:"sections"`
}
