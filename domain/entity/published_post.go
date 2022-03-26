package entity

import (
	"gopkg.in/guregu/null.v4"
)

type PublishedPost struct {
	ID int64 `json:"id" db:"id"`
}

type SearchResultArticles struct {
	NextUrl null.String `json:"next_url"`
	Data    []PostList  `json:"data"`
}

type PostList struct {
	ID          int64       `json:"-"`
	Title       string      `json:"title"`
	Slug        string      `json:"slug"`
	Excerpt     null.String `json:"excerpt"`
	ArticleUrl  string      `json:"article_url"`
	ReleaseDate int64       `json:"release_date"`
	IsCSC       bool        `json:"is_csc"`

	PostTypeID   int64 `json:"-"`
	CategoryID   int64 `json:"-"`
	CreatorID    int64 `json:"-"`
	CoverMediaID int64 `json:"-"`

	PostType *SearchResultPostType `json:"post_type"`
	Category *SearchResultCategory `json:"category"`
	Cover    *Cover                `json:"cover"`

	Author  []Author `json:"author"`
	Creator *User    `json:"-"`

	PublishAt null.String `json:"-"`
}
