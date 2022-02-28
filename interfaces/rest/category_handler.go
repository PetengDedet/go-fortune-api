package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/gin-gonic/gin"
)

type CategoryHandler struct {
	CategoryApp application.CategoryApp
}

func NewCategoryHandler(categoryApp application.CategoryApp) *CategoryHandler {
	return &CategoryHandler{
		CategoryApp: categoryApp,
	}
}

func (handler *CategoryHandler) GetCategoryPageDetailHandler(c *gin.Context) {
	slug := c.Param("categorySlug")
	category, err := handler.CategoryApp.GetCategoryPageDetailBySlug(slug)
	if err != nil {
		if errors.Is(err, &common.NotFoundError{}) {
			c.JSON(http.StatusNotFound, NotFoundResponse(nil))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(category))
}
