package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type CategoryHandler struct {
	CategoryApp application.CategoryApp
	PageApp     application.PageApp
}

func NewCategoryHandler(categoryApp application.CategoryApp, pageApp application.PageApp) *CategoryHandler {
	return &CategoryHandler{
		CategoryApp: categoryApp,
		PageApp:     pageApp,
	}
}

func (handler *CategoryHandler) GetCategoryPageDetailHandler(c echo.Context) error {
	slug := c.Param("categorySlug")
	category, err := handler.CategoryApp.GetCategoryPageDetailBySlug(slug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	categoryPage, err := handler.PageApp.GetCategoryPageDetail("category", category)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(categoryPage))
}
