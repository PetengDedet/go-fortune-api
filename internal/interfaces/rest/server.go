package rest

import (
	"net/http"
	"os"

	"github.com/PetengDedet/fortune-post-api/internal/application"
	"github.com/PetengDedet/fortune-post-api/internal/infrastructure/persistence/mongodb"
	"github.com/PetengDedet/fortune-post-api/internal/infrastructure/persistence/mysql"
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	postRepo := mysql.PostRepo{
		DB: db,
	}
	userRepo := mysql.UserRepo{
		DB: db,
	}
	postDetailRepo := mysql.PostDetailRepo{
		DB: db,
	}
	mediaRepo := mysql.MediaRepo{
		DB: db,
	}
	magazineRepo := mysql.MagazineRepo{
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
		PostRepo:         &postRepo,
		UserRepo:         &userRepo,
		CategoryRepo:     &categoryRepo,
		TagRepo:          &tagRepo,
		PostDetailRepo:   &postDetailRepo,
		MediaRepo:        &mediaRepo,
		LinkoutRepo:      &linkoutRepo,
		PostTypeRepo:     &postTypeRepo,
	}
	keywordMongoApp := application.KeywordApp{
		KeywordRepo: &keywordMongoRepo,
	}
	keywordApp := application.KeywordApp{
		KeywordRepo: &keywordRepo,
	}
	magazineApp := application.MagazineApp{
		MagazineRepo: &magazineRepo,
	}

	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"app":    "Fortune API",
		})
	})

	v1 := e.Group("/v1")
	v1.GET("/menu", NewMenuHandler(menuApp).GetPublicMenuPositionsHandler)
	v1.GET("/category/:categorySlug", NewCategoryHandler(categoryApp, pageApp).GetCategoryPageDetailHandler)
	v1.GET("/tag/:tagSlug", NewTagHandler(tagApp, pageApp).GetTagPageDetailHandler)
	v1.GET("/search", NewSearchHandler(searchApp, pageApp).GetSearchResultHandler)
	v1.POST("/search", NewKeywordHandler(&keywordMongoApp).SaveKeywordHandler)
	v1.GET("/popular-keyword", NewKeywordHandler(&keywordApp).GetPopularKeywordHandler)
	v1.GET("/content-type/:postTypeSlug", NewPostTypeHandler(postTypeApp, pageApp).GetPostTypePageHandler)
	v1.GET("/most-popular", NewPublishedPostHandler(publishedPostApp).GetMostPopularPostHandler)
	v1.GET("/related-articles", NewPublishedPostHandler(publishedPostApp).GetRelatedArticlesHandler)
	v1.GET("/magazines", NewMagazineHandler(magazineApp).GetLatestHomepageMagazines)
	v1.GET("/latest", NewPublishedPostHandler(publishedPostApp).GetLatestArticleHandler)
	v1.GET("/latest/homepage/tag/:tagSlug", NewPublishedPostHandler(publishedPostApp).GetLatestArticleHomepageByTagHandler)
	v1.GET("/latest/homepage/content-type/:contentTypeSlug", NewPublishedPostHandler(publishedPostApp).GetLatestArticleHomepageByContentTypeHandler)
	v1.GET("/latest/tag/:tagSlug", NewPublishedPostHandler(publishedPostApp).GetLatestArticleByTagHandler)
	v1.GET("/latest/content-type/:contentTypeSlug", NewPublishedPostHandler(publishedPostApp).GetLatestArticleByContentTypeHandler)
	v1.GET("/latest/category/:categorySlug", NewPublishedPostHandler(publishedPostApp).GetLatestArticleByCategoryHandler)

	v1.GET("/:pageSlug", NewPageHandler(pageApp).GetPageBySlugHandler)
	v1.GET("/:categorySlug/amp/:authorUsername/:postSlug", NewPublishedPostHandler(publishedPostApp).GetAMPDetailArticleHandler)
	v1.GET("/:categorySlug/:authorUsername/:postSlug", NewPublishedPostHandler(publishedPostApp).GetDetailArticleHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
