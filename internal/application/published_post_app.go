package application

import (
	"os"
	"strconv"
	"time"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
	"gopkg.in/guregu/null.v4"
)

type PublishedPostApp struct {
	PublishePostRepo repository.PublishedPostRepository
	PostRepo         repository.PostRepository
	UserRepo         repository.UserRepository
	TagRepo          repository.TagRepository
	CategoryRepo     repository.CategoryRepository
	PostDetailRepo   repository.PostDetailRepository
	MediaRepo        repository.MediaRepository
	LinkoutRepo      repository.LinkoutRepository
	PostTypeRepo     repository.PostTypeRepository
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

func (app *PublishedPostApp) GetPostDetails(categorySlug, authorUsername, postSlug string) (*entity.PublishedPost, error) {
	post, err := app.PublishePostRepo.GetPublishedPostDetail(categorySlug, authorUsername, postSlug)
	if err != nil {
		return nil, err
	}

	tags, err := app.TagRepo.GetTagsByPostId(post.ID)
	if err != nil {
		return nil, err
	}
	post.Tags = tags

	authors, err := app.UserRepo.GetAuthorsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	for i, a := range authors {
		authors[i] = *a.SetAvatar()
	}

	post.Author = authors

	postDetail, err := app.PostDetailRepo.GetPostDetailsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	linkouts, err := app.LinkoutRepo.GetLinkoutsByType("article")
	if err != nil {
		return nil, err
	}

	for i, pd := range postDetail {
		if pd.Type.String != "cover" && pd.Type.String != "image" {
			postDetail[i].Value.String = entity.ParseDescription(pd.Value.String, post.Title, linkouts)
			continue
		}

		postDetail[i].Value = null.String{}
	}

	post.Description.String = entity.ParseDescription(post.Description.String, post.Title, linkouts)
	post.PostDetails = postDetail

	// Update visited count
	now := time.Now()
	if err := app.PublishePostRepo.IncrementVisitCount(post.ID, &now); err != nil {
		panic(err)
	}

	if err := app.PostRepo.IncrementVisitCount(post.ID, &now); err != nil {
		panic(err)
	}

	return post, nil
}

func (app *PublishedPostApp) GetAMPPostDetails(categorySlug, authorUsername, postSlug string) (*entity.PublishedPost, error) {
	post, err := app.PublishePostRepo.GetPublishedPostDetail(categorySlug, authorUsername, postSlug)
	if err != nil {
		return nil, err
	}

	tags, err := app.TagRepo.GetTagsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	if len(tags) <= 0 {
		tags = []entity.Tag{}
	}

	post.Tags = tags

	authors, err := app.UserRepo.GetAuthorsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	for i, a := range authors {
		authors[i] = *a.SetAvatar()
	}

	post.Author = authors

	postDetail, err := app.PostDetailRepo.GetPostDetailsByPostId(post.ID)
	if err != nil {
		return nil, err
	}

	linkouts, err := app.LinkoutRepo.GetLinkoutsByType("article")
	if err != nil {
		return nil, err
	}

	for i, pd := range postDetail {
		if pd.Type.String != "cover" && pd.Type.String != "image" {
			val := entity.ParseDescription(pd.Value.String, post.Title, linkouts)
			val = entity.AdjustAMPDescription(val)

			postDetail[i].Value.String = val
			postDetail[i].Embeds = entity.GetAMPEmbedsType(val)

			continue
		}

		postDetail[i].Value = null.String{}
	}

	parsedDescription := entity.ParseDescription(post.Description.String, post.Title, linkouts)
	post.Description.String = entity.AdjustAMPDescription(parsedDescription)
	post.PostDetails = postDetail

	// Update visited count
	now := time.Now()
	if err := app.PublishePostRepo.IncrementVisitCount(post.ID, &now); err != nil {
		panic(err)
	}

	if err := app.PostRepo.IncrementVisitCount(post.ID, &now); err != nil {
		panic(err)
	}

	return post, nil
}

func (app *PublishedPostApp) GetLatestPost(page int) (pl []entity.PostList, e error) {
	perPage1, err := strconv.Atoi(os.Getenv("LATEST_HOMEPAGE_MAIN_PER_PAGE_1"))
	if err != nil {
		perPage1 = 12
	}

	perPage2, err := strconv.Atoi(os.Getenv("LATEST_HOMEPAGE_MAIN_PER_PAGE_2"))
	if err != nil {
		perPage2 = 9
	}

	maxPage, err := strconv.Atoi(os.Getenv("LATEST_HOMEPAGE_MAIN_MAX_PAGE"))
	if err != nil {
		maxPage = 10
	}

	if page == 0 {
		page = 1
	}

	if page > maxPage {
		page = maxPage
	}

	take := perPage1
	skip := 0

	if page > 1 && perPage2 > 0 {
		skip = take + ((page - 2) * perPage2)
		take = perPage2
	}

	posts, err := app.PublishePostRepo.GetLatestPublishedPost(take, skip)
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
	pl = posts

	return pl, nil
}

func (app *PublishedPostApp) GetLatestPostHomepageByTagSLug(tagSlug string) (pl []entity.PostList, e error) {
	take, err := strconv.Atoi(os.Getenv("LATEST_HOMEPAGE_TAG_PER_PAGE"))
	if err != nil {
		take = 6
	}

	tag, err := app.TagRepo.GetTagBySlug(tagSlug)
	if err != nil {
		return nil, err
	}

	posts, err := app.PublishePostRepo.GetLatestPublishedPostByTagId(take, 0, tag.ID)
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
	pl = posts

	return pl, nil
}

func (app *PublishedPostApp) GetLatestPostByContentTypeSLug(contentTypeSlug string) (pl []entity.PostList, e error) {
	take, err := strconv.Atoi(os.Getenv("LATEST_HOMEPAGE_CONTENT_TYPE_PER_PAGE"))
	if err != nil {
		take = 6
	}

	contentType, err := app.PostTypeRepo.GetPostTypeBySlug(contentTypeSlug)
	if err != nil {
		return nil, err
	}

	posts, err := app.PublishePostRepo.GetLatestPublishedPostByPostTypeId(take, 0, contentType.ID)
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
	pl = posts

	return pl, nil
}

func (app *PublishedPostApp) GetLatestPostByTagSLug(tagSlug string, page int) (pl []entity.PostList, e error) {
	perPage1, err := strconv.Atoi(os.Getenv("LATEST_TAG_PER_PAGE_1"))
	if err != nil {
		perPage1 = 14
	}

	perPage2, err := strconv.Atoi(os.Getenv("LATEST_TAG_PER_PAGE_2"))
	if err != nil {
		perPage2 = 9
	}

	maxPage, err := strconv.Atoi(os.Getenv("LATEST_TAG_MAX_PAGE"))
	if err != nil {
		maxPage = 10
	}

	if page == 0 {
		page = 1
	}

	if page > maxPage {
		page = maxPage
	}

	take := perPage1
	skip := 0

	if page > 1 && perPage2 > 0 {
		skip = take + ((page - 2) * perPage2)
		take = perPage2
	}

	tag, err := app.TagRepo.GetTagBySlug(tagSlug)
	if err != nil {
		return nil, err
	}

	posts, err := app.PublishePostRepo.GetLatestPublishedPostByTagId(take, skip, tag.ID)
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
	pl = posts

	return pl, nil
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
