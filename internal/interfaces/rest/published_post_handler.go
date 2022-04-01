package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type PublishedPostHandlerInterface interface {
	GetMostPopularPostHandler()
}

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
