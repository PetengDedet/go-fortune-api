package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/gin-gonic/gin"
)

type PageHandler struct {
	PageApp application.PageApp
}

func NewPageHandler(pageApp application.PageApp) *PageHandler {
	return &PageHandler{
		PageApp: pageApp,
	}
}

func (pageHandler *PageHandler) GetPageBySlugHandler(c *gin.Context) {
	slug := c.Param("pageSlug")
	pageDetail, err := pageHandler.PageApp.GetPageDetailBySlug(slug)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			c.JSON(http.StatusNotFound, NotFoundResponse(nil))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(pageDetail))
}
