package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type TagRepository interface {
	GetTagByIds(ids []int64) ([]entity.Tag, error)
	GetTagBySlug(slug string) (*entity.Tag, error)
	GetPostIdsByTagId(id int64) ([]int64, error)
	GetTagsByPostId(postId int64) ([]entity.Tag, error)
}
