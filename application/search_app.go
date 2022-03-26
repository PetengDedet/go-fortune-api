package application

import (
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type SearchApp struct {
	PublishedPostRepo repository.PublishedPostRepository
	UserRepo          repository.UserRepository
}

func (app *SearchApp) GetSearchResult(keyword string, page int) ([]entity.PostList, error) {
	limit, err := strconv.Atoi(os.Getenv("SEARCH_RESULT_LIMIT"))
	if err != nil || limit <= 0 {
		limit = 5
	}

	if page <= 0 {
		page = 1
	}
	skip := limit * (page - 1)

	posts := []entity.PostList{}
	if len(keyword) <= 0 {
		latestPosts, err := app.PublishedPostRepo.GetLatestPublishedPost(limit, skip)
		if err != nil {
			return nil, err
		}

		posts = latestPosts
	}

	if len(keyword) > 0 {
		searchResult, err := app.PublishedPostRepo.SearchPublishedPostByKeyword(keyword, limit, skip)
		if err != nil {
			return nil, err
		}

		posts = searchResult
	}

	if len(posts) == 0 {
		return []entity.PostList{}, nil
	}

	var postIds []int64
	for _, p := range posts {
		postIds = append(postIds, p.ID)
	}

	authors, err := app.UserRepo.GetAuthorsByPostIds(postIds)
	posts = mapAuthorToPost(posts, authors)

	return posts, nil
}
