package entity

import "gopkg.in/guregu/null.v4"

type Page struct {
	ID              int64       `json:"-"`
	Name            null.String `json:"-"`
	Page            null.String `json:"page"`
	Slug            null.String `json:"slug"`
	Excerpt         null.String `json:"excerpt"`
	ArticleCounts   int64       `json:"article_counts"`
	Description     null.String `json:"description"`
	Url             null.String `json:"url"`
	MetaTitle       null.String `json:"meta_title"`
	MetaDescription null.String `json:"meta_description"`
	Sections        []Section   `json:"sections"`
}

func PageResponse(page *Page) *Page {
	page.Page = page.Name

	return page
}
