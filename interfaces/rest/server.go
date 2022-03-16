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

	// Repos
	menuRepo := mysql.MenuRepo{
		DB: db,
	}
	pageRepo := mysql.PageRepo{
		DB: db,
	}
	sectionRepo := mysql.SectionRepo{
		DB: db,
	}
	categoryRepo := mysql.CategoryRepo{
		DB: db,
	}
	rankRepo := mysql.RankRepo{
		DB: db,
	}
	rankCategoryRepo := mysql.RankCategoryRepo{
		DB: db,
	}
	linkoutRepo := mysql.LinkoutRepo{
		DB: db,
	}
	tagRepo := mysql.TagRepo{
		DB: db,
	}
	postTypeRepo := mysql.PostTypeRepo{
		DB: db,
	}

	// Apps
	menuApp := application.MenuApp{
		MenuRepo:     &menuRepo,
		PageRepo:     &pageRepo,
		CategoryRepo: &categoryRepo,
		RankRepo:     &rankRepo,
		LinkoutRepo:  &linkoutRepo,
	}
	pageApp := application.PageApp{
		PageRepo:         &pageRepo,
		SectionRepo:      &sectionRepo,
		PostTypeRepo:     &postTypeRepo,
		TagRepo:          &tagRepo,
		RankRepo:         &rankRepo,
		RankCategoryRepo: &rankCategoryRepo,
		CategoryRepo:     &categoryRepo,
		LinkoutRepo:      &linkoutRepo,
	}

	// categoryApp := application.CategoryApp{
	// 	CategoryRepo: &categoryRepo,
	// 	SectionRepo:  &sectionRepo,
	// }

	v1 := route.Group("/v1")
	{
		v1.GET("/menu", NewMenuHandler(menuApp).GetPublicMenuPositionsHandler)
		v1.GET("/:pageSlug", NewPageHandler(pageApp).GetPageBySlugHandler)
		// v1.GET("/category/:categorySlug", NewCategoryHandler(categoryApp).GetCategoryPageDetailHandler)

		latest := route.Group("/latest")
		{
			latest.GET("/", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"route": "/v1/latest",
				})
			})
			latest.GET("/homepage/tag/:slug", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"route": "/v1/latest/homepage/tag/" + c.Param("slug"),
				})
			})
			latest.GET("/homepage/content-type/:slug", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"route": "/v1/latest/homepage/content-type/" + c.Param("slug"),
				})
			})
			// test.GET("/tag/:slug", func (c *gin.Context)  {
			// 	c.JSON(http.StatusOK, gin.H{
			// 		"route": "/v1/latest/homepage/content-type/" + c.Param("slug"),
			// 	})
			// })

		}
	}

	route.Run(":8000")
}
