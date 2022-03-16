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
	Width       int64       `json:"width" db:"width"`
	Height      int64       `json:"height" db:"height"`
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

func (m *Media) SetUrl(width, height int) *Media {
	cdnDomain := os.Getenv("CDN_DOMAIN")
	if cdnDomain == "" {
		cdnDomain = "https://cdn.fortuneidn.com/"
	}

	fileBaseName := strings.TrimSuffix(filepath.Base(m.UrlMedia.String), filepath.Ext(m.UrlMedia.String))
	extension := filepath.Ext(m.UrlMedia.String)
	dirName := filepath.Dir(m.UrlMedia.String)

	m.Url = null.StringFrom(fmt.Sprintf("%s%s/%s_%dx%d%s", cdnDomain, dirName, fileBaseName, width, height, extension))

	return m
}
