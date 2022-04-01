package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type PostDetailRepository interface {
	GetPostDetailsByPostId(postId int64) ([]entity.PostDetailList, error)
}
