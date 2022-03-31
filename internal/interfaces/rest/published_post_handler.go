package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/gin-gonic/gin"
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

func (handler PublishedPostHandler) GetMostPopularPostHandler(c *gin.Context) {
	posts, err := handler.PublishedPostApp.GetMostPopularPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(posts))
}

func (handler PublishedPostHandler) GetRelatedArticlesHandler(c *gin.Context) {
	page := c.Query("page")
	categorySlug := c.Query("categorySlug")
	tagSlug := c.Query("tagSlug")

	postList, err := handler.PublishedPostApp.GeRelatedPosts(page, tagSlug, categorySlug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			c.JSON(http.StatusNotFound, NotFoundResponse(nil))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(postList))
}
