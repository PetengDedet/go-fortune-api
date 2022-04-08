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
	description = adjustDescriptionIframeSize(description)

	domDesc, err := dom.FastParse(strings.NewReader(description))
	if err != nil {
		panic(err)
	}

	domDesc = adjustDescriptionImageAlt(domDesc, title)
	domDesc = adjustEmbedToBeResponsive(domDesc)
	domDesc = adjustImageSourceUrl(domDesc)
	domDesc = adjustLinks(domDesc, linkouts)

	description = strings.TrimPrefix(dom.InnerHTML(domDesc), "<html><head></head><body>")
	description = strings.TrimSuffix(description, "</body></html>")
	description = strings.ReplaceAll(description, "\n", " ")

	return description
}

func adjustDescriptionIframeSize(description string) string {
	return strings.Replace(description, "<div class=\"video-container\"><iframe", "<div class=\"embed-image lazyYT\"><iframe width=\"100%\" height=\"420\"", -1)
}

func adjustDescriptionImageAlt(node *html.Node, title string) *html.Node {
	title = strings.Replace(title, "\"", "", -1)

	imgs := dom.GetAllNodesWithTag(node, "img")
	for _, v := range imgs {
		dom.SetAttribute(v, "alt", title)
	}

	return node
}

func adjustEmbedToBeResponsive(node *html.Node) *html.Node {
	embeds := dom.QuerySelectorAll(node, "div[data-oembed-url]")
	for _, v := range embeds {
		dom.SetAttribute(v, "class", "embeddedContent")
	}

	return node
}

func adjustImageSourceUrl(node *html.Node) *html.Node {
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
func adjustLinks(node *html.Node, whiteLists []Linkout) *html.Node {
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

func AdjustAMPDescription(description string) string {
	domDesc, err := dom.FastParse(strings.NewReader(description))
	if err != nil {
		panic(err)
	}

	domDesc = adjustAMPInstagram(domDesc)
	domDesc = adjustAMPFacebook(domDesc)
	domDesc = adjustAMPTiktok(domDesc)
	domDesc = adjustAMPTwitter(domDesc)
	domDesc = adjustAMPYoutube(domDesc)
	domDesc = adjustAMPIframe(domDesc)

	description = strings.TrimPrefix(dom.InnerHTML(domDesc), "<html><head></head><body>")
	description = strings.TrimSuffix(description, "</body></html>")
	description = strings.ReplaceAll(description, "\n", " ")

	return description
}

func GetAMPEmbedsType(description string) (embeds []string) {
	node, err := dom.FastParse(strings.NewReader(description))
	if err != nil {
		panic(err)
	}

	if dom.QuerySelector(node, "amp-instagram") != nil {
		embeds = append(embeds, "instagram")
	}

	if dom.QuerySelector(node, "amp-facebook") != nil {
		embeds = append(embeds, "facebook")
	}

	if dom.QuerySelector(node, "amp-tiktok") != nil {
		embeds = append(embeds, "tiktok")
	}

	if dom.QuerySelector(node, "amp-twitter") != nil {
		embeds = append(embeds, "twitter")
	}

	if dom.QuerySelector(node, "amp-youtube") != nil {
		embeds = append(embeds, "youtube")
	}

	if dom.QuerySelector(node, "amp-iframe") != nil {
		embeds = append(embeds, "iframe")
	}

	return
}

func adjustAMPInstagram(node *html.Node) *html.Node {
	igEmbeds := dom.QuerySelectorAll(node, ".instagram-media")
	var embeds []*html.Node
	var shortCodes []string

	for _, n := range igEmbeds {
		var re = regexp.MustCompile(`instagram.com\/p\/(.*?)\/`)
		parent := n.Parent
		str := dom.GetAttribute(parent, "data-oembed-url")
		match := re.FindStringSubmatch(str)
		if len(match) > 0 {
			embeds = append(embeds, parent)
			shortCodes = append(shortCodes, match[1])
		}
	}

	for i, n := range embeds {
		ampIg := dom.CreateElement("amp-instagram")
		dom.SetAttribute(ampIg, "data-shortcode", shortCodes[i])
		dom.SetAttribute(ampIg, "width", "250")
		dom.SetAttribute(ampIg, "height", "250")
		dom.SetAttribute(ampIg, "layout", "responsive")
		dom.SetAttribute(ampIg, "data-captioned", "true")

		dom.ReplaceChild(n.Parent, ampIg, n)
	}

	return node
}

func adjustAMPFacebook(node *html.Node) *html.Node {
	fbEmbeds := dom.QuerySelectorAll(node, ".fb-post")
	var embeds []*html.Node
	var links []string

	for _, i := range fbEmbeds {
		parent := i.Parent
		fbLink := dom.GetAttribute(parent, "data-oembed-url")
		embeds = append(embeds, parent)
		links = append(links, fbLink)
	}

	for i, n := range embeds {
		ampFb := dom.CreateElement("amp-facebook")
		dom.SetAttribute(ampFb, "data-href", links[i])
		dom.SetAttribute(ampFb, "width", "250")
		dom.SetAttribute(ampFb, "height", "510")
		dom.SetAttribute(ampFb, "layout", "responsive")
		dom.SetAttribute(ampFb, "data-align-center", "true")
		if strings.Contains(links[i], "/videos/") {
			dom.SetAttribute(ampFb, "data-embed-as", "video")
		}

		dom.ReplaceChild(n.Parent, ampFb, n)
	}

	return node
}

func adjustAMPTiktok(node *html.Node) *html.Node {
	tkEmbeds := dom.QuerySelectorAll(node, ".tiktok-embed")
	var embeds []*html.Node
	var videoIds []string

	for _, n := range tkEmbeds {
		videoId := dom.GetAttribute(n, "data-video-id")
		if len(videoId) > 0 {
			embeds = append(embeds, n.Parent)
			videoIds = append(videoIds, videoId)
		}
	}

	for i, n := range embeds {
		wrapper := dom.CreateElement("div")
		dom.SetAttribute(wrapper, "style", "display: flex;justify-content: center;align-items: center;")

		ampTk := dom.CreateElement("amp-tiktok")
		dom.SetAttribute(ampTk, "data-src", videoIds[i])
		dom.SetAttribute(ampTk, "width", "325")
		dom.SetAttribute(ampTk, "height", "575")

		dom.AppendChild(wrapper, ampTk)

		dom.ReplaceChild(n.Parent, wrapper, n)
	}

	return node
}

func adjustAMPTwitter(node *html.Node) *html.Node {
	twEmbeds := dom.QuerySelectorAll(node, ".twitter-tweet")
	var embeds []*html.Node
	var tweetIds []string

	for _, i := range twEmbeds {
		parent := i.Parent
		var re = regexp.MustCompile(`\/status\/(\d+)`)
		str := dom.GetAttribute(parent, "data-oembed-url")
		match := re.FindStringSubmatch(str)
		if len(match) > 0 {
			embeds = append(embeds, parent)
			tweetIds = append(tweetIds, match[1])
		}
	}

	for i, n := range embeds {
		ampTw := dom.CreateElement("amp-twitter")
		dom.SetAttribute(ampTw, "data-tweetid", tweetIds[i])
		dom.SetAttribute(ampTw, "width", "250")
		dom.SetAttribute(ampTw, "height", "315")
		dom.SetAttribute(ampTw, "layout", "responsive")

		dom.ReplaceChild(n.Parent, ampTw, n)
	}

	return node
}

func adjustAMPYoutube(node *html.Node) *html.Node {
	ytEmbeds := dom.QuerySelectorAll(node, ".embeddedContent")
	var embeds []*html.Node
	var ytIds []string

	for _, i := range ytEmbeds {
		str := dom.GetAttribute(i, "data-oembed-url")
		ytId := getYoutubeId(str)
		if ytId != "" {
			embeds = append(embeds, i)
			ytIds = append(ytIds, ytId)
		}
	}

	for i, n := range embeds {
		ampYt := dom.CreateElement("amp-youtube")
		dom.SetAttribute(ampYt, "data-videoid", ytIds[i])
		dom.SetAttribute(ampYt, "width", "480")
		dom.SetAttribute(ampYt, "height", "270")
		dom.SetAttribute(ampYt, "layout", "responsive")

		dom.ReplaceChild(n.Parent, ampYt, n)
	}

	return node
}

func getYoutubeId(str string) (link string) {
	patterns := []string{
		`youtube.com\/embed\/(\w+)`,
		`youtu.be\/(\w+)`,
		`youtube.com\/watch\?v\=(\w+)`,
	}

	for _, p := range patterns {
		var re = regexp.MustCompile(p)
		match := re.FindStringSubmatch(str)
		if len(match) > 0 {
			link = match[1]
			break
		}
	}

	return
}

func adjustAMPIframe(node *html.Node) *html.Node {
	ifEmbeds := dom.QuerySelectorAll(node, ".embeddedContent")
	var embeds []*html.Node
	var urls []string

	for _, i := range ifEmbeds {
		url := dom.GetAttribute(i, "data-oembed-url")
		if len(url) > 0 {
			embeds = append(embeds, i)
			urls = append(urls, url)
		}
	}

	for i, n := range embeds {
		ampIf := dom.CreateElement("amp-iframe")
		dom.SetAttribute(ampIf, "src", urls[i])
		dom.SetAttribute(ampIf, "width", "350")
		dom.SetAttribute(ampIf, "height", "220")
		dom.SetAttribute(ampIf, "layout", "responsive")
		dom.SetAttribute(ampIf, "resizable", "true")
		dom.SetAttribute(ampIf, "frameborder", "0")
		dom.SetAttribute(ampIf, "sandbox", "allow-scripts allow-same-origin")

		dom.ReplaceChild(n.Parent, ampIf, n)
	}

	return node
}
