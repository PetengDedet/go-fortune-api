package application

import (
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type SearchApp struct {
	PageRepo          repository.PageRepository
	PublishedPostRepo repository.PublishedPostRepository
}

func (sa *SearchApp) GetSearchResult(keyword string, page int) ([]entity.SearchResultArticle, error) {
	limit, err := strconv.Atoi(os.Getenv("SEARCH_RESULT_LIMIT"))
	if err != nil || limit <= 0 {
		limit = 5
	}

	if page <= 0 {
		page = 1
	}

	skip := limit * (page - 1)
	limit = limit * page

	// TODO: handle empty keyword

	articleResult, err := sa.PublishedPostRepo.SearchPublishedPostByKeyword(keyword, limit, skip)
	if err != nil {
		return nil, err
	}

	// TODO: handle empty result
	var postIds []int
	for _, post := range articleResult {
		postIds = append(postIds, int(post.ID))
	}

	authors, err := sa.PublishedPostRepo.GetAuthorsByPostIds(postIds)
	if err != nil {
		panic(err)
	}

	articleResult = mapAuthotToPost(articleResult, authors)

	return articleResult, nil
}

func mapAuthotToPost(posts []entity.SearchResultArticle, authors []entity.Author) []entity.SearchResultArticle {
	for i, p := range posts {
		for _, a := range authors {
			if a.PostID == p.ID {
				posts[i].Author = append(posts[i].Author, *a.SetAvatar())
			}
		}
	}

	return posts
}
