package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
)

type MagazineHandler struct {
	MagazineApp application.MagazineApp
}

func NewMagazineHandler(magazineApp application.MagazineApp) *MagazineHandler {
	return &MagazineHandler{
		MagazineApp: magazineApp,
	}
}

func (handler *MagazineHandler) GetLatestHomepageMagazines(c echo.Context) error {
	magazines, err := handler.MagazineApp.GetLatestMagazines()
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(magazines))
}
