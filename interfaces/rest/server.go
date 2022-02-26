package rest

import (
	"net/http"
	"os"

	"github.com/PetengDedet/fortune-post-api/application"
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

	db, err := mysql.GetConnection(db_host, db_port, db_name, db_username, db_password)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	menuRepo := mysql.MenuRepo{
		DB: db,
	}
	menuApp := application.MenuApp{
		MenuRepo: menuRepo,
	}

	v1 := route.Group("/v1")
	{
		v1.GET("/menu", func(c *gin.Context) {
			publicMenuPosition := menuApp.GetPublicMenuPositions()

			c.JSON(http.StatusOK, gin.H{
				"data":    publicMenuPosition,
				"status":  http.StatusOK,
				"message": "success",
			})
		})
	}

	route.Run(":8000")
}
