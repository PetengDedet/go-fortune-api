package rest

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/common"
	"github.com/labstack/echo/v4"
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

func (handler *SearchHandler) GetSearchResultHandler(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	keyword = strings.TrimSpace(keyword)

	pageParam := c.QueryParam("page")
	if len(pageParam) <= 0 {
		pageParam = "1"
	}

	currentPage, err := strconv.Atoi(pageParam)
	if err != nil {
		currentPage = 1
	}

	searchResult, err := handler.SearchApp.GetSearchResult(keyword, currentPage)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	searchPage, err := handler.PageApp.GetSearchResultPageDetail(keyword, currentPage, searchResult)
	if err != nil {
		if errors.Is(&common.NotFoundError{}, err) {
			return c.JSON(http.StatusNotFound, NotFoundResponse(nil))
		}

		return c.JSON(http.StatusInternalServerError, InternalErrorResponse(nil))
	}

	return c.JSON(http.StatusOK, SuccessResponse(searchPage))
}
