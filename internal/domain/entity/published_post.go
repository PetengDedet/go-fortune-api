package entity

import (
	"os"
	"regexp"
	"strings"

	"github.com/go-shiori/dom"
	"golang.org/x/net/html"
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

	PostType *PostTypeList `json:"post_type"`
	Category *CategoryList `json:"category"`
	Cover    *Cover        `json:"cover"`

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

func ParseDescription(description, title string, linkouts []Linkout) string {
	description = AdjustDescriptionIframeSize(description)

	domDesc, err := dom.FastParse(strings.NewReader(description))
	if err != nil {
		panic(err)
	}

	domDesc = AdjustDescriptionImageAlt(domDesc, title)
	domDesc = AdjustEmbedToBeResponsive(domDesc)
	domDesc = AdjustImageSourceUrl(domDesc)
	domDesc = AdjustLinks(domDesc, linkouts)

	description = strings.TrimPrefix(dom.InnerHTML(domDesc), "<html><head></head><body>")
	description = strings.TrimSuffix(description, "</body></html>")
	description = strings.ReplaceAll(description, "\n", " ")

	return description
}

func AdjustDescriptionIframeSize(description string) string {
	return strings.Replace(description, "<div class=\"video-container\"><iframe", "<div class=\"embed-image lazyYT\"><iframe width=\"100%\" height=\"420\"", -1)
}

func AdjustDescriptionImageAlt(node *html.Node, title string) *html.Node {
	title = strings.Replace(title, "\"", "", -1)

	imgs := dom.GetAllNodesWithTag(node, "img")
	for _, v := range imgs {
		dom.SetAttribute(v, "alt", title)
	}

	return node
}

func AdjustEmbedToBeResponsive(node *html.Node) *html.Node {
	embeds := dom.QuerySelectorAll(node, "div[data-oembed-url]")
	for _, v := range embeds {
		dom.SetAttribute(v, "class", "embeddedContent")
	}

	return node
}

func AdjustImageSourceUrl(node *html.Node) *html.Node {
	imgSources := dom.QuerySelectorAll(node, "span.main-article-source")
	for _, v := range imgSources {
		link := dom.QuerySelector(v, "a")
		source := dom.InnerHTML(dom.FirstElementChild(link))
		dom.SetInnerHTML(v, source)
	}

	return node
}

// If the link match the whitelist pattern,
// then keep the link but add a target="_blank"
// Otherwise, remove the link, but keep the text
func AdjustLinks(node *html.Node, whiteLists []Linkout) *html.Node {
	links := dom.GetAllNodesWithTag(node, "a")

	// No need to check
	if len(links) == 0 {
		return node
	}

	domain := os.Getenv("DOMAIN")
	if len(domain) <= 0 {
		domain = "fortuneidn.com"
	}

	var patterns []string
	patterns = append(patterns, `http(.*)`+domain)
	for _, p := range whiteLists {
		ps := strings.TrimPrefix(p.Url, "/")
		ps = strings.TrimSuffix(ps, "/U")
		patterns = append(patterns, ps)
	}

	var allowedLinks []*html.Node
	var disAllowedLinks []*html.Node

	for _, l := range links {
		isAllowed := false
		for _, p := range patterns {
			var re = regexp.MustCompile(p)
			href := dom.GetAttribute(l, "href")
			if len(re.FindStringIndex(href)) > 0 {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			allowedLinks = append(allowedLinks, l)
			continue
		}

		disAllowedLinks = append(disAllowedLinks, l)
	}

	for _, l := range allowedLinks {
		dom.SetAttribute(l, "target", "_blank")
	}

	for _, l := range disAllowedLinks {
		parent := l.Parent
		text := dom.CreateTextNode(dom.TextContent(l))
		dom.ReplaceChild(parent, text, l)
	}

	return node
}
