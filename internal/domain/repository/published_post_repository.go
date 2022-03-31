package repository

import "github.com/PetengDedet/fortune-post-api/internal/domain/entity"

type PublishedPostRepository interface {
	GetPublishedPostCountByCategoryId(catId int64) (int64, error)
	GetPublishedPostCountByTagId(tagId int64) (int64, error)
	GetPublishedPostCountByPostTypeId(postTypeId int64) (int64, error)
	SearchPublishedPostByKeyword(keyword string, limit, skip int) ([]entity.PostList, error)
	GetLatestPublishedPost(limit, skip int) ([]entity.PostList, error)
	GetLatestPublishedPostByCategoryId(limit, skip int, categoryId int64) ([]entity.PostList, error)
	GetLatestPublishedPostByTagId(limit, skip int, tagId int64) ([]entity.PostList, error)
	GetLatestPublishedPostByCategoryIdAndTagId(limit, skip int, categoryId, tagId int64) ([]entity.PostList, error)
	GetPopularPosts() ([]entity.PostList, error)
}
