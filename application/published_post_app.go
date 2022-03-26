package application

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
)

type PublishedPostAppInterface interface {
	GetMostPopularPosts() ([]entity.PostList, error)
}

type PublishedPostApp struct {
	PublishePostRepo repository.PublishedPostRepository
	UserRepo         repository.UserRepository
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
