package entity

type Category struct {
	ID              int64
	GeneralStatusID int64
	Name            string
	Slug            string
	Excerpt         string
	MediaID         int64
	MetaTitle       string
	MetaDescription string
}
