package repository

type PublishedPostRepository interface {
	GetPublishedPostCountByCategoryId(catId int64) (int64, error)
}
