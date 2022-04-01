package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type TagHandler struct {
	TagApp  application.TagApp
	PageApp application.PageApp
}

func NewTagHandler(tagApp application.TagApp, pageApp application.PageApp) *TagHandler {
	return &TagHandler{
		TagApp:  tagApp,
		PageApp: pageApp,
	}
}

func (handler *TagHandler) GetTagPageDetailHandler(c echo.Context) error {
	slug := c.Param("tagSlug")
	tag, err := handler.TagApp.GetTagDetail(slug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	tagPage, err := handler.PageApp.GetTagPageDetail("tag", tag)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(tagPage))
}
