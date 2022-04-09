package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type PublishedPostHandler struct {
	PublishedPostApp application.PublishedPostApp
}

func NewPublishedPostHandler(app application.PublishedPostApp) *PublishedPostHandler {
	return &PublishedPostHandler{
		PublishedPostApp: app,
	}
}

func (handler PublishedPostHandler) GetMostPopularPostHandler(c echo.Context) error {
	posts, err := handler.PublishedPostApp.GetMostPopularPosts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(posts))
}

func (handler PublishedPostHandler) GetRelatedArticlesHandler(c echo.Context) error {
	page := c.QueryParam("page")
	categorySlug := c.QueryParam("categorySlug")
	tagSlug := c.QueryParam("tagSlug")

	postList, err := handler.PublishedPostApp.GeRelatedPosts(page, tagSlug, categorySlug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(postList))
}

func (handler *PublishedPostHandler) GetDetailArticleHandler(c echo.Context) error {
	categorySlug := c.Param("categorySlug")
	authorUsername := c.Param("authorUsername")
	postSlug := c.Param("postSlug")

	post, err := handler.PublishedPostApp.GetPostDetails(categorySlug, authorUsername, postSlug)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, SuccessResponse(post))
}

func (handler *PublishedPostHandler) GetAMPDetailArticleHandler(c echo.Context) error {
	categorySlug := c.Param("categorySlug")
	authorUsername := c.Param("authorUsername")
	postSlug := c.Param("postSlug")

	post, err := handler.PublishedPostApp.GetAMPPostDetails(categorySlug, authorUsername, postSlug)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, SuccessResponse(post))
}

func (handler *PublishedPostHandler) GetLatestArticleHandler(c echo.Context) error {
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 1
	}

	postList, err := handler.PublishedPostApp.GetLatestPost(page)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(postList))
}

func (handler *PublishedPostHandler) GetLatestArticleByTagHandler(c echo.Context) error {
	tagSlug := c.Param("tagSlug")
	postList, err := handler.PublishedPostApp.GetLatestPostByTagSLug(tagSlug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(postList))
}
