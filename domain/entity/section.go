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
	Slug                 null.String `json:"-"`
	MediaID              null.String `json:"-"`
	PageSectionID        null.Int    `json:"-"`
	PageSectionPageID    null.Int    `json:"-"`
	PageSectionSectionID null.Int    `json:"-"`
	CategoryName         null.String `json:"-"`
	CategorySlug         null.String `json:"-"`
	TagName              null.String `json:"-"`
	TagSlug              null.String `json:"-"`
	PostTypeName         null.String `json:"-"`
	PostTypeSlug         null.String `json:"-"`
	LinkoutUrl           null.String `json:"-"`
	LinkoutType          null.String `json:"-"`
	RankName             null.String `json:"-"`
	RankSlug             null.String `json:"-"`
	RankCategoryName     null.String `json:"-"`
	RankCategorySlug     null.String `json:"-"`
	UserUsername         null.String `json:"-"`
}

func SectionResponse(s *Section) *Section {
	s.BaseUrl = generateBaseUrl(s)
	s.Url = generateUrl(s)
	// s.Url = generatePublicUrl(s)

	return s
}

func generateBaseUrl(s *Section) null.String {
	if s.TableName.String == "categories" {
		return null.StringFrom("/" + s.CategorySlug.String)
	}

	if s.TableName.String == "post_types" {
		return null.StringFrom("/" + s.PostTypeSlug.String)
	}

	if s.TableName.String == "tags" {
		return null.StringFrom("/tag/" + s.TagSlug.String)
	}

	if s.TableName.String == "users" {
		return null.StringFrom("/" + s.UserUsername.String)
	}

	return s.BaseUrl
}

func generateUrl(s *Section) null.String {
	url := null.StringFrom("/v1/" + s.Type.String)

	if s.Type.String == "latest-home" {
		url = null.StringFrom("/v1/latest")
		return url
	}

	if s.TableName.String != "" {
		if s.TableName.String != "categories" {
			url = null.StringFrom(url.String + "/homepage/")
		}

		if s.TableName.String == "categories" {
			url = null.StringFrom("/v1/" + os.Getenv("CATEGORY_SECTION_TYPE"))
			url = null.StringFrom(url.String + "?category=" + s.CategorySlug.String)
			return url
		}

		if s.TableName.String == "post_types" {
			url = null.StringFrom(url.String + "content-type/" + s.PostTypeSlug.String)
			return url
		}

		if s.TableName.String == "tags" {
			url = null.StringFrom(url.String + "tag/" + s.TagSlug.String)
			return url
		}

		if s.TableName.String == "users" {
			url = null.StringFrom(url.String + "author/" + s.UserUsername.String)
			return url
		}

		if s.TableName.String == "linkouts" {
			url = s.LinkoutUrl
			return url
		}
	}

	return url
}
