package rest

import (
	"net/http"
	"os"

	"github.com/PetengDedet/fortune-post-api/infrastructure/persistence/mysql"
	"github.com/gin-gonic/gin"
)

func Init() {
	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"app":    "Fortune API",
		})
	})

	db_host := os.Getenv("MYSQL_HOST")
	db_port := os.Getenv("MYSQL_PORT")
	db_name := os.Getenv("MYSQL_DBNAME")
	db_username := os.Getenv("MYSQL_USERNAME")
	db_password := os.Getenv("MYSQL_PASSWORD")

	services, err := mysql.NewRepositories(db_host, db_port, db_name, db_username, db_password)
	if err != nil {
		panic(err)
	}
	defer services.Close()

	menus := NewMenuHandler(services.MenuRepo)
	route.GET("/menus", menus.GetMenuPositions)

	//
	route.Run(":8000")
}
