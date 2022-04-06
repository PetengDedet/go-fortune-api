package entity

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/guregu/null.v4"
)

type Media struct {
	ID          null.Int    `json:"id"`
	Name        string      `json:"name" db:"name"`
	Description null.String `json:"descrtiption" db:"description"`
	UrlMedia    null.String `json:"url_media" db:"url_media"`
	Mime        null.String `json:"mime" db:"mime"`
	Extension   null.String `json:"extension" db:"extension"`
	SourceName  null.String `json:"source_name" db:"source_name"`
	SourceUrl   null.String `json:"source_url" db:"source_url"`
	Width       int16       `json:"width" db:"width"`
	Height      int16       `json:"height" db:"height"`
	Keyword     null.String `json:"keyword" db:"keyword"`
	UrlEmbed    null.String `json:"url_embed" db:"url_embed"`
	Url         null.String `json:"url" db:"-"`

	GeneralStatusID int64     `json:"-" db:"general_status_id"`
	GeneralTypeID   int64     `json:"-" db:"general_type_id"`
	GalleryID       null.Int  `json:"-" db:"gallery_id"`
	CreatedAt       null.Time `json:"-" db:"created_at"`
	UpdatedAt       null.Time `json:"-" db:"updated_at"`
	CreatedBy       int64     `json:"-" db:"created_by"`
	UpdatedBy       null.Int  `json:"-" db:"updated_by"`

	GeneralType   *GeneralType   `json:"-" db:"-"`
	GeneralStatus *GeneralStatus `json:"-" db:"-"`
	Gallery       *Gallery       `json:"-" db:"-"`
	Creator       *User          `json:"-" db:"-"`
	Updater       *User          `json:"-" db:"-"`
}

func (m *Media) SetUrl(width, height string) *Media {
	m.Url = getUrl(m.UrlMedia, width, height)

	return m
}

func getUrl(urlMedia null.String, width, height string) (url null.String) {
	cdnDomain := os.Getenv("CDN_DOMAIN")
	if cdnDomain == "" {
		cdnDomain = "https://cdn.fortuneidn.com/"
	}

	fileBaseName := strings.TrimSuffix(filepath.Base(urlMedia.String), filepath.Ext(urlMedia.String))
	extension := filepath.Ext(urlMedia.String)
	dirName := filepath.Dir(urlMedia.String)

	return null.StringFrom(fmt.Sprintf("%s%s/%s_%sx%s%s", cdnDomain, dirName, fileBaseName, width, height, extension))
}

type Cover struct {
	Full        null.String `json:"full"`
	Large       null.String `json:"large"`
	Medium      null.String `json:"medium"`
	Thumbnail   null.String `json:"thumbnail"`
	Small       null.String `json:"small"`
	Tiny        null.String `json:"tiny"`
	SourceName  null.String `json:"source_name"`
	SourceUrl   null.String `json:"source_url"`
	Description null.String `json:"description"`
	Width       null.Int    `json:"width"`
	Height      null.Int    `json:"height"`
	EmbedVideo  null.String `json:"embed_video"`

	UrlMedia null.String `json:"-"`
}

func (c *Cover) GetPredefinedSize() *Cover {
	c.Full = getUrl(c.UrlMedia, "1050", "700")
	c.Large = getUrl(c.UrlMedia, "1000", "auto")
	c.Medium = getUrl(c.UrlMedia, "600", "auto")
	c.Small = getUrl(c.UrlMedia, "200", "auto")
	c.Thumbnail = getUrl(c.UrlMedia, "600", "400")
	c.Tiny = getUrl(c.UrlMedia, "6", "4")

	return c
}
