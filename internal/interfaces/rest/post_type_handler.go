package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
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

func (handler *PostTypeHandler) GetPostTypePageHandler(c echo.Context) error {
	slug := c.Param("postTypeSlug")
	if len(slug) == 0 {
		return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
	}

	postType, err := handler.PostTypeApp.GetPostTypeDetail(slug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, nil)
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	page, err := handler.PageApp.GetPostTypePageDetail(postType)
	if err != nil {
		if err != nil {
			if errors.Is(err, &common.NotFoundError{}) {
				return c.JSON(http.StatusNotFound, nil)
			}

			return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		}
	}

	return c.JSON(http.StatusOK, SuccessResponse(page))
}
