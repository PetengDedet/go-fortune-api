package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/application"
	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/gin-gonic/gin"
)

type SearchHandler struct {
	SearchApp application.SearchApp
	PageApp   application.PageApp
}

func NewSearchHandler(searchApp application.SearchApp, pageApp application.PageApp) *SearchHandler {
	return &SearchHandler{
		SearchApp: searchApp,
		PageApp:   pageApp,
	}
}

func (handler *SearchHandler) GetSearchResultHandler(c *gin.Context) {
	keyword, ok := c.GetQuery("keyword")
	if !ok {
		keyword = ""
		// TODO: get latest article
	}

	pageParam, ok := c.GetQuery("page")
	if !ok {
		pageParam = "1"
	}

	currentPage, err := strconv.Atoi(pageParam)
	if err != nil {
		currentPage = 1
	}

	searchResult, err := handler.SearchApp.GetSearchResult(keyword, currentPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	searchPage, err := handler.PageApp.GetSearchResultPageDetail(keyword, currentPage, searchResult)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			c.JSON(http.StatusNotFound, NotFoundResponse(nil))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
		return
	}

	c.JSON(http.StatusOK, SuccessResponse(searchPage))
}
