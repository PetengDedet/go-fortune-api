package entity

import (
	"gopkg.in/guregu/null.v4"
)

type Section struct {
	ID       int64       `json:"-"`
	Type     null.String `json:"type"`
	Title    null.String `json:"title"`
	OrderNum int64       `json:"order_num"`
	Url      null.String `json:"url"`
	BaseUrl  null.String `json:"base_url"`
	ImageUrl null.String `json:"image_url"`

	TableName            null.String `json:"-"`
	TableID              null.Int    `json:"-"`
	TableUrl             null.String `json:"-"`
	Slug                 null.String `json:"-"`
	MediaID              null.String `json:"-"`
	PageSectionID        null.Int    `json:"-"`
	PageSectionPageID    null.Int    `json:"-"`
	PageSectionSectionID null.Int    `json:"-"`

	Category     *Category     `json:"-" db:"-"`
	Linkout      *Linkout      `json:"-" db:"-"`
	Rank         *Rank         `json:"-" db:"-"`
	RankCategory *RankCategory `json:"-" db:"-"`
	PostType     *PostType     `json:"-" db:"-"`
	Tag          *Tag          `json:"-" db:"-"`

	Media *Media `json:"-" db:"-"`
}

func (s *Section) SetSectionAttributes() *Section {
	s.setUrl()
	s.setBaseUrl()

	if s.Media != nil {
		s.ImageUrl = s.Media.SetUrl("600", "600").Url
	}

	return s
}

func (s *Section) setBaseUrl() *Section {

	if s.Type.String == "latest" && s.Tag != nil {
		s.BaseUrl = null.StringFrom("/tag/" + s.Tag.Slug)
	}

	if s.Type.String == "category" && s.Category != nil {
		s.BaseUrl = null.StringFrom("/" + s.Category.Slug)
	}

	return s
}

func (s *Section) setUrl() *Section {

	if s.Type.String == "headline" {
		s.Url = null.StringFrom("/v1/headline")
	}

	if s.Type.String == "headline" && s.TableName.String == "categories" && s.Category == nil {
		s.Url = null.StringFrom("/v1/headline?category=")
	}

	if s.Type.String == "latest-home" {
		s.Url = null.StringFrom("/v1/latest")
	}

	if s.Type.String == "latest" && s.Tag != nil {
		s.Url = null.StringFrom("/v1/latest/homepage/tag/" + s.Tag.Slug)
	}

	if s.Type.String == "latest" && s.TableName.String == "tags" && s.Tag == nil {
		s.Url = null.StringFrom("/v1/latest/tag/")
	}

	if s.Type.String == "latest" && s.TableName.String == "categories" && s.Category == nil {
		s.Url = null.StringFrom("/v1/latest/category/")
	}

	if s.Type.String == "category" && s.Category != nil {
		s.Url = null.StringFrom("/v1/headline?category=" + s.Category.Slug)
	}

	if s.Type.String == "magazines" {
		s.Url = null.StringFrom("/v1/magazines")
	}

	if s.Type.String == "most-popular" {
		s.Url = null.StringFrom("/v1/most-popular")
	}

	if s.Type.String == "popular-keyword" {
		s.Url = null.StringFrom("/v1/popular-keyword")
	}

	if s.Type.String == "search" {
		s.Url = null.StringFrom("/v1/search")
	}

	if s.Type.String == "rank-category" {
		s.Url = null.StringFrom("/v1/rank-category")
	}

	if s.Linkout != nil {
		s.Url = null.StringFrom(s.Linkout.Url)
	}

	if s.Type.String == "rank-awardee" {
		s.Url = null.StringFrom("/v1/rank-awardee")
	}

	return s
}

func (s *Section) MutateUrl(url string) *Section {
	s.Url = null.StringFrom(url)

	return s
}

func (s *Section) MutateBaseUrl(baseUrl string) *Section {
	s.BaseUrl = null.StringFrom(baseUrl)

	return s
}
