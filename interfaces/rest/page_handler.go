package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "failed",
			"data":    nil,
			"error": gin.H{
				"message":          "Something went wrong",
				"reason":           "internal_error",
				"error_user_title": nil,
				"error_user_msg":   "Something went wrong",
			},
		})
		return
	}

	if pageDetail == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "failed",
			"data":    nil,
			"error": gin.H{
				"message":          "Page doesn't exist",
				"reason":           "not_found",
				"error_user_title": nil,
				"error_user_msg":   "Page doesn't exist",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success",
		"data":    pageDetail,
	})
}
