package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/labstack/echo/v4"
)

type MenuHandler struct {
	MenuApp application.MenuApp
}

func NewMenuHandler(menuApp application.MenuApp) *MenuHandler {
	return &MenuHandler{
		MenuApp: menuApp,
	}
}

func (menuHandler *MenuHandler) GetPublicMenuPositionsHandler(c echo.Context) error {
	publicMenuPosition, err := menuHandler.MenuApp.GetPublicMenuPositions()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(publicMenuPosition))
}
