package repository

type KeywordRepository interface {
	SaveNewKeyword(keyword string) error
}
