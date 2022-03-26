package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
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
