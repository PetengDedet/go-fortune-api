package repository

type PostRepository interface {
	GetPostCountByCategoryId(catId int64) (int64, error)
}
