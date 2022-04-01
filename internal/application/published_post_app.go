package application

import (
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type PublishedPostApp struct {
	PublishePostRepo repository.PublishedPostRepository
	UserRepo         repository.UserRepository
	TagRepo          repository.TagRepository
	CategoryRepo     repository.CategoryRepository
	PostDetailRepo   repository.PostDetailRepository
	MediaRepo        repository.MediaRepository
}

func (app *PublishedPostApp) GetMostPopularPosts() ([]entity.PostList, error) {
	posts, err := app.PublishePostRepo.GetPopularPosts()
	if err != nil {
		return nil, err
	}

	if len(posts) == 0 {
		return []entity.PostList{}, nil
	}

	var postIds []int64
	for _, p := range posts {
		postIds = append(postIds, p.ID)
	}

	authors, err := app.UserRepo.GetAuthorsByPostIds(postIds)
	if err != nil {
		panic(err)
	}

	posts = mapAuthorToPost(posts, authors)

	return posts, nil
}

func (app *PublishedPostApp) GeRelatedPosts(page, tagSlug, categorySlug string) (pl []entity.PostList, e error) {
	limit, err := strconv.Atoi(os.Getenv("RELATED_ARTICLES_LIMIT"))
	if err != nil {
		limit = 6
	}

	currentPage, err := strconv.Atoi(page)
	if err != nil {
		currentPage = 1
	}

	skip := limit * (currentPage - 1)

	// Tag and category
	if len(tagSlug) > 0 && len(categorySlug) > 0 {
		tag, err := app.TagRepo.GetTagBySlug(tagSlug)
		if err != nil {
			return nil, err
		}

		cat, err := app.CategoryRepo.GetCategoryBySlug(categorySlug)
		if err != nil {
			return nil, err
		}

		pl, e = app.PublishePostRepo.GetLatestPublishedPostByCategoryIdAndTagId(limit, skip, cat.ID.Int64, tag.ID)
		if e != nil {
			return nil, e
		}
	}

	// Tag only
	if len(tagSlug) > 0 && len(categorySlug) == 0 {
		tag, err := app.TagRepo.GetTagBySlug(tagSlug)
		if err != nil {
			return nil, err
		}

		pl, e = app.PublishePostRepo.GetLatestPublishedPostByTagId(limit, skip, tag.ID)
		if e != nil {
			return nil, e
		}
	}

	// Category only
	if len(tagSlug) == 0 && len(categorySlug) > 0 {
		cat, err := app.CategoryRepo.GetCategoryBySlug(categorySlug)
		if err != nil {
			return nil, err
		}

		pl, e = app.PublishePostRepo.GetLatestPublishedPostByCategoryId(limit, skip, cat.ID.Int64)
		if e != nil {
			return nil, e
		}
	}

	// No tag neither category
	if len(tagSlug) == 0 && len(categorySlug) == 0 {
		pl, e = app.PublishePostRepo.GetLatestPublishedPost(limit, skip)
		if e != nil {
			return nil, e
		}
	}

	if len(pl) == 0 {
		return []entity.PostList{}, nil
	}

	postIds := []int64{}
	for _, p := range pl {
		postIds = append(postIds, p.ID)
	}

	authors, err := app.UserRepo.GetAuthorsByPostIds(postIds)
	if err != nil {
		panic(err)
	}

	pl = mapAuthorToPost(pl, authors)

	return pl, e
}

func (app *PublishedPostApp) GetPostDetails(categorySlug, authorSlug, postSlug string) (*entity.PublishedPost, error) {

	return nil, nil
}

func mapAuthorToPost(posts []entity.PostList, authors []entity.Author) []entity.PostList {
	for i, p := range posts {
		for _, a := range authors {
			if a.PostID == p.ID {
				posts[i].Author = append(posts[i].Author, *a.SetAvatar())
			}
		}
	}

	return posts
}
