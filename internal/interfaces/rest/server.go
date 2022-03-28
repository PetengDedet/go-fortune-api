package rest

import (
	"net/http"
	"os"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/infrastructure/persistence/mongodb"
	"github.com/PetengDedet/fortune-post-api/internal/infrastructure/persistence/mysql"
	"github.com/gin-gonic/gin"
)

func Init() {
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

	mongoClient := mongodb.GetMongoClient(os.Getenv("MONGODB_URI"))
	mongoDB := mongodb.GetDB(mongoClient, os.Getenv("MONGODB_DATABASE"))
	defer mongodb.CloseMongoConnection()
	// Repos
	keywordMongoRepo := mongodb.KeywordRepo{
		DB: mongoDB,
	}

	keywordRepo := mysql.KeywordRepo{
		DB: db,
	}

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
	publishedPostRepo := mysql.PublishedPostRepo{
		DB: db,
	}
	userRepo := mysql.UserRepo{
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
	categoryApp := application.CategoryApp{
		CategoryRepo:      &categoryRepo,
		SectionRepo:       &sectionRepo,
		PublishedPostRepo: &publishedPostRepo,
	}
	tagApp := application.TagApp{
		TagRepo:           &tagRepo,
		SectionRepo:       &sectionRepo,
		PublishedPostRepo: &publishedPostRepo,
	}
	searchApp := application.SearchApp{
		PublishedPostRepo: &publishedPostRepo,
		UserRepo:          &userRepo,
	}
	postTypeApp := application.PostTypeApp{
		PostTypeRepo:      &postTypeRepo,
		PublishedPostRepo: &publishedPostRepo,
	}
	publishedPostApp := application.PublishedPostApp{
		PublishePostRepo: &publishedPostRepo,
		UserRepo:         &userRepo,
	}
	keywordMongoApp := application.KeywordApp{
		KeywordRepo: &keywordMongoRepo,
	}
	keywordApp := application.KeywordApp{
		KeywordRepo: &keywordRepo,
	}

	route := gin.Default()

	route.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"app":    "Fortune API",
		})
	})

	v1 := route.Group("/v1")
	{
		v1.GET("/menu", NewMenuHandler(menuApp).GetPublicMenuPositionsHandler)
		v1.GET("/category/:categorySlug", NewCategoryHandler(categoryApp, pageApp).GetCategoryPageDetailHandler)
		v1.GET("/tag/:tagSlug", NewTagHandler(tagApp, pageApp).GetTagPageDetailHandler)
		v1.GET("/search", NewSearchHandler(searchApp, pageApp).GetSearchResultHandler)
		v1.POST("/search", NewKeywordHandler(&keywordMongoApp).SaveKeywordHandler)
		v1.GET("/popular-keyword", NewKeywordHandler(&keywordApp).GetPopularKeywordHandler)
		v1.GET("/content-type/:postTypeSlug", NewPostTypeHandler(postTypeApp, pageApp).GetPostTypePageHandler)
		v1.GET("/most-popular", NewPublishedPostHandler(publishedPostApp).GetMostPopularPostHandler)

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
		}

		v1.GET("/:pageSlug", NewPageHandler(pageApp).GetPageBySlugHandler)
	}

	route.Run(":8000")
}
