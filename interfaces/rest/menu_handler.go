package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
	"github.com/gin-gonic/gin"
)

type MenuHandler struct {
	MenuApp application.MenuApp
}

func NewMenuHandler(menuApp application.MenuApp) *MenuHandler {
	return &MenuHandler{
		MenuApp: menuApp,
	}
}

func (menuHandler *MenuHandler) GetPublicMenuPositionsHandler(c *gin.Context) {
	publicMenuPosition, err := menuHandler.MenuApp.GetPublicMenuPositions()
	if err != nil {

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(publicMenuPosition))
}
