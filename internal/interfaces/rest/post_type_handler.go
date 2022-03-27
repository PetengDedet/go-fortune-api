package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/gin-gonic/gin"
)

type PostTypeHandler struct {
	PostTypeApp application.PostTypeApp
	PageApp     application.PageApp
}

func NewPostTypeHandler(postTypeApp application.PostTypeApp, pageApp application.PageApp) *PostTypeHandler {
	return &PostTypeHandler{
		PostTypeApp: postTypeApp,
		PageApp:     pageApp,
	}
}

func (handler *PostTypeHandler) GetPostTypePageHandler(c *gin.Context) {
	slug := c.Param("postTypeSlug")
	if len(slug) == 0 {
		c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		return
	}

	postType, err := handler.PostTypeApp.GetPostTypeDetail(slug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			c.JSON(http.StatusNotFound, nil)
			return
		}

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	page, err := handler.PageApp.GetPostTypePageDetail(postType)
	if err != nil {
		if err != nil {
			if errors.Is(err, &common.NotFoundError{}) {
				c.JSON(http.StatusNotFound, nil)
				return
			}

			c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
			return
		}
	}

	c.JSON(http.StatusOK, SuccessResponse(page))
}
