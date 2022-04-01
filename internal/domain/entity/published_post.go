package entity

import (
	"gopkg.in/guregu/null.v4"
)

type PublishedPost struct {
	ID int64 `json:"-" db:"id"`

	Title           string      `json:"title" db:"title"`
	Slug            string      `json:"slug" db:"slug"`
	Excerpt         null.String `json:"excerpt" db:"excerpt"`
	Description     null.String `json:"description" db:"description"`
	IsCSC           bool        `json:"is_csc" db:"is_csc"`
	MetaTitle       null.String `json:"meta_title" db:"meta_title"`
	MetaDescription null.String `json:"meta_description" db:"meta_description"`
	OGCaption       null.String `json:"og_caption" db:"og_caption"`
	OGTitle         null.String `json:"og_title" db:"og_title"`
	OGDescription   null.String `json:"og_description" db:"og_desctiption"`

	ArticleUrl  string    `json:"article_url" db:"-"`
	ReleaseDate int64     `json:"release_date" db:"-"`
	Campaign    *Campaign `json:"campaign" db:"-"`

	PostTypeID   int64 `json:"-"`
	CategoryID   int64 `json:"-"`
	CreatorID    int64 `json:"-"`
	CoverMediaID int64 `json:"-"`

	PostType *PostType     `json:"post_type"`
	Category *CategoryList `json:"category"`
	Cover    *CoverDetail  `json:"cover"`

	Author      []Author         `json:"author"`
	Tags        []Tag            `json:"tags"`
	PostDetails []PostDetailList `json:"post_details"`
	Creator     *User            `json:"-"`

	PublishAt null.String `json:"-"`
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

	PostType *PostTypeList `json:"post_type"`
	Category *CategoryList `json:"category"`
	Cover    *Cover        `json:"cover"`

	Author  []Author `json:"author"`
	Creator *User    `json:"-"`

	PublishAt null.String `json:"-"`
}
