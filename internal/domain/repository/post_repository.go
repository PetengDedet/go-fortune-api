package repository

import "time"

type PostRepository interface {
	GetPostCountByCategoryId(catId int64) (int64, error)
	IncrementVisitCount(postId int64, updatedAt *time.Time) error
}
