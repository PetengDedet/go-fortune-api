package rest

import (
	"net/http"

	"github.com/PetengDedet/fortune-post-api/application"
	"github.com/gin-gonic/gin"
)

//Menuss struct defines the dependencies that will be used
type MenuHandler struct {
	app application.MenuAppInterface
}

//Users constructor
func NewMenuHandler(app application.MenuAppInterface) *MenuHandler {
	return &MenuHandler{
		app: app,
	}
}

func (handler *MenuHandler) GetMenuPositions(c *gin.Context) {
	// menuPositions := []entity.MenuPosition{} //customize user
	// var err error
	// menuPositions, err = handler.app.GetMenuPositions()
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// fmt.Println(menuPositions)

	c.JSON(http.StatusOK, nil)
}
