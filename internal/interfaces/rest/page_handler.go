package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type PageHandler struct {
	PageApp application.PageApp
}

func NewPageHandler(pageApp application.PageApp) *PageHandler {
	return &PageHandler{
		PageApp: pageApp,
	}
}

func (pageHandler *PageHandler) GetPageBySlugHandler(c echo.Context) error {
	slug := c.Param("pageSlug")
	pageDetail, err := pageHandler.PageApp.GetPageDetailBySlug(slug)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(pageDetail))
}
