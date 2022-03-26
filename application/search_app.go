package application

import (
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type SearchApp struct {
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

	articles := []entity.SearchResultArticle{}
	if len(keyword) <= 0 {
		latestArticle, err := sa.PublishedPostRepo.GetLatestPublishedPost(limit, skip)
		if err != nil {
			return nil, err
		}

		articles = latestArticle
	}

	if len(keyword) > 0 {
		searchResult, err := sa.PublishedPostRepo.SearchPublishedPostByKeyword(keyword, limit, skip)
		if err != nil {
			return nil, err
		}

		articles = searchResult
	}

	if len(articles) == 0 {
		return []entity.SearchResultArticle{}, nil
	}

	authors := getAuthors(sa, articles)
	articles = mapAuthotToPost(articles, authors)

	return articles, nil
}

func getAuthors(sa *SearchApp, articles []entity.SearchResultArticle) (authors []entity.Author) {
	if len(articles) <= 0 {
		return
	}

	var postIds []int
	for _, post := range articles {
		postIds = append(postIds, int(post.ID))
	}

	newAuthors, err := sa.PublishedPostRepo.GetAuthorsByPostIds(postIds)
	if err != nil {
		panic(err)
	}

	return newAuthors
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
