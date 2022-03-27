package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/gin-gonic/gin"
)

type KeywordHandler struct {
	KeywordApp application.KeywordApp
}

func NewKeywordHandler(keywordApp *application.KeywordApp) *KeywordHandler {
	return &KeywordHandler{
		KeywordApp: *keywordApp,
	}
}

func (handler *KeywordHandler) SaveKeywordHandler(c *gin.Context) {
	kw := &entity.KeywordHistory{}

	if err := c.ShouldBindJSON(&kw); err != nil {
		c.JSON(http.StatusBadRequest, BadRequestResponse("keyword needed", "keyword needed"))
		return
	}

	err := handler.KeywordApp.SaveKeyword(kw.Keyword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(gin.H{}))
}
