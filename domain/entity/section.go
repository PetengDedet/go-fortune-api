package entity

import (
	"os"

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
	CategoryName         null.String `json:"-"`
	CategorySlug         null.String `json:"-"`
	CategoryUrl          null.String `json:"-"`
	TagName              null.String `json:"-"`
	TagSlug              null.String `json:"-"`
	TagUrl               null.String `json:"-"`
	PostTypeName         null.String `json:"-"`
	PostTypeSlug         null.String `json:"-"`
	LinkoutUrl           null.String `json:"-"`
	LinkoutType          null.String `json:"-"`
	RankName             null.String `json:"-"`
	RankSlug             null.String `json:"-"`
	RankCategoryName     null.String `json:"-"`
	RankCategorySlug     null.String `json:"-"`
	UserUsername         null.String `json:"-"`
	SourceTableSlug      null.String `json:"-"`
}

func SectionResponse(s *Section) *Section {
	s.SourceTableSlug = getSourceTableSlugAttribute(s)
	s.Type = getTypeResponse(s)
	s.Url = getUrlAttribute(s)
	// s.Url = getUrlResponse(s)
	s.BaseUrl = getBaseUrlAttribute(s)
	// s.BaseUrl = getBaseUrlRespose(s)

	return s
}

func getBaseUrlAttribute(s *Section) null.String {
	url := null.StringFrom("")

	if s.TableName.String != "" {
		if s.TableName.String == "categories" {
			url = null.StringFrom("/" + s.CategorySlug.String)
		}

		if s.TableName.String == "post_types" {
			url = null.StringFrom("/" + s.PostTypeSlug.String)
		}

		if s.TableName.String == "tags" {
			url = null.StringFrom("/tag/" + s.TagSlug.String)
		}

		if s.TableName.String == "users" {
			url = null.StringFrom("/" + s.UserUsername.String)
		}
	}

	return url
}

func getUrlAttribute(s *Section) null.String {
	url := null.StringFrom("/v1/" + s.Type.String)

	if s.Type.String == "latest-home" {
		url = null.StringFrom("/v1/latest")
	}

	if s.TableName.String != "" {
		if s.TableName.String != "categories" {
			url = null.StringFrom(url.String + "/homepage/")
		}

		switch s.TableName.String {
		case "categories":
			url = null.StringFrom("/v1/" + os.Getenv("CATEGORY_SECTION_TYPE"))
			url = null.StringFrom(url.String + "?category=" + s.CategorySlug.String)

		case "post_types":
			url = null.StringFrom(url.String + "content-type/" + s.PostTypeSlug.String)

		case "tags":
			url = null.StringFrom(url.String + "tag/" + s.TagSlug.String)

		case "users":
			url = null.StringFrom(url.String + "author/" + s.UserUsername.String)

		case "linkouts":
			url = s.LinkoutUrl
		default:
		}
	}

	return url
}

func getTypeResponse(s *Section) null.String {
	t := s.Type
	if s.TableName.String == "post_types" {
		t = null.StringFrom(t.String + "-" + s.SourceTableSlug.String)
	}

	return t
}

func getSourceTableSlugAttribute(s *Section) null.String {
	ts := null.StringFrom("")
	if s.TableName.String == "categories" {
		return s.CategorySlug
	}

	if s.TableName.String == "post_types" {
		return s.PostTypeSlug
	}

	if s.TableName.String == "tags" {
		return s.TagSlug
	}

	if s.TableName.String == "users" {
		return s.UserUsername
	}

	if s.TableName.String == "linkouts" {
		return s.LinkoutUrl
	}

	return ts
}
