package rest

import (
	"errors"
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type KeywordHandler struct {
	KeywordApp application.KeywordApp
}

func NewKeywordHandler(keywordApp *application.KeywordApp) *KeywordHandler {
	return &KeywordHandler{
		KeywordApp: *keywordApp,
	}
}

func (handler *KeywordHandler) SaveKeywordHandler(c echo.Context) error {
	// kw := &entity.KeywordHistory{}

	// if err := c.bi ShouldBindJSON(&kw); err != nil {
	// 	return c.JSON(http.StatusBadRequest, BadRequestResponse("keyword needed", "keyword needed"))
	// }

	// err := handler.KeywordApp.SaveKeyword(kw.Keyword)
	// if err != nil {
	// 	return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	// }

	return c.JSON(http.StatusOK, SuccessResponse(gin.H{}))
}

func (handler *KeywordHandler) GetPopularKeywordHandler(c echo.Context) error {
	kw, err := handler.KeywordApp.GetPopularKeyword()
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(kw))
}
