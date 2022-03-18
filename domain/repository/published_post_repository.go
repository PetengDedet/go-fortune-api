package repository

type PublishedPostRepository interface {
	GetPublishedPostCountByCategoryId(catId int64) (int64, error)
	GetPublishedPostCountByTagId(tagId int64) (int64, error)
}
