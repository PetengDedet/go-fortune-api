package mysql

import (
	"strings"
	"time"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PublishedPostRepo struct {
	DB *sqlx.DB
}

func (ppr *PublishedPostRepo) GetPublishedPostCountByCategoryId(catId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts
		WHERE category_id = ?
	`

	err := ppr.DB.Get(&postCount, query, catId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (ppr *PublishedPostRepo) GetPublishedPostCountByTagId(tagId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts pp
		INNER JOIN post_tags pt ON pt.post_id = pp.id 
		WHERE pt.tag_id = ?
	`

	err := ppr.DB.Get(&postCount, query, tagId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (ppr *PublishedPostRepo) GetPublishedPostCountByPostTypeId(postTypeId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts pp
		WHERE pp.post_type_id = ?
	`

	err := ppr.DB.Get(&postCount, query, postTypeId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (ppr *PublishedPostRepo) GetLatestPublishedPost(limit, skip int) ([]entity.PostList, error) {
	query := `
		SELECT
			pp.id, 
			pp.title,
			pp.slug,
			pp.publish_at,
			pp.is_csc,
			pp.post_type_id,
			pp.category_id,
			pp.cover_media_id,
			pp.created_by,
			m.url_media,
			c.name,
			c.slug,
			pt.name,
			pt.slug,
			u.username
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
		ORDER BY pp.publish_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := ppr.DB.Query(query, limit, skip)
	if err != nil {
		return nil, err
	}

	var sr []entity.PostList
	for rows.Next() {
		var art entity.PostList
		var cover entity.Cover
		var srCat entity.SearchResultCategory
		var srPt entity.SearchResultPostType
		var publishAt string
		var username *string
		err := rows.Scan(
			&art.ID,
			&art.Title,
			&art.Slug,
			&publishAt,
			&art.IsCSC,
			&art.PostTypeID,
			&art.CategoryID,
			&art.CoverMediaID,
			&art.CreatorID,
			&cover.UrlMedia,
			&srCat.Name,
			&srCat.Slug,
			&srPt.Name,
			&srPt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		srCat.Url = "/" + srCat.Slug
		art.Category = &srCat
		art.Cover = cover.GetPredefinedSize()
		art.PostType = &srPt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		art.ReleaseDate = t.Unix()

		art.ArticleUrl = "/" + srCat.Slug + "/" + *username + "/" + art.Slug

		sr = append(sr, art)
	}

	return sr, nil
}

func (ppr *PublishedPostRepo) GetPopularPosts() ([]entity.PostList, error) {
	query := `
		SELECT
			pp.id, 
			pp.title,
			pp.slug,
			pp.publish_at,
			pp.is_csc,
			pp.post_type_id,
			pp.category_id,
			pp.cover_media_id,
			pp.created_by,
			pp.excerpt,
			m.url_media,
			c.name,
			c.slug,
			pt.name,
			pt.slug,
			u.username
		FROM post_populars ppo
			INNER JOIN published_posts pp ON pp.id = ppo.post_id
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
		ORDER BY ppo.order_num
	`

	rows, err := ppr.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var sr []entity.PostList
	for rows.Next() {
		var art entity.PostList
		var cover entity.Cover
		var srCat entity.SearchResultCategory
		var srPt entity.SearchResultPostType
		var publishAt string
		var username *string
		err := rows.Scan(
			&art.ID,
			&art.Title,
			&art.Slug,
			&publishAt,
			&art.IsCSC,
			&art.PostTypeID,
			&art.CategoryID,
			&art.CoverMediaID,
			&art.CreatorID,
			&art.Excerpt,
			&cover.UrlMedia,
			&srCat.Name,
			&srCat.Slug,
			&srPt.Name,
			&srPt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		srCat.Url = "/" + srCat.Slug
		art.Category = &srCat
		art.Cover = cover.GetPredefinedSize()
		art.PostType = &srPt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		art.ReleaseDate = t.Unix()

		art.ArticleUrl = "/" + srCat.Slug + "/" + *username + "/" + art.Slug

		sr = append(sr, art)
	}

	return sr, nil
}

func (ppr *PublishedPostRepo) SearchPublishedPostByKeyword(keyword string, limit, skip int) ([]entity.PostList, error) {
	relevant, lessRelevant := formatKeyword(keyword)
	stmt, err := ppr.DB.Preparex(`
		SELECT
			pp.id, 
			pp.title,
			pp.slug,
			pp.publish_at,
			pp.is_csc,
			pp.post_type_id,
			pp.category_id,
			pp.cover_media_id,
			pp.created_by,
			m.url_media,
			c.name,
			c.slug,
			pt.name,
			pt.slug,
			u.username,
			MATCH (title) AGAINST (? IN BOOLEAN MODE) AS relevant,
			MATCH (title) AGAINST (? IN BOOLEAN MODE) AS less_relevant
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
		WHERE MATCH (pp.title) AGAINST (? IN BOOLEAN MODE)
		ORDER BY relevant DESC, less_relevant DESC
		LIMIT ? OFFSET ?
	`)

	if err != nil {
		panic(err)
	}

	rows, err := stmt.Query(relevant, lessRelevant, lessRelevant, limit, skip)
	if err != nil {
		return nil, err
	}

	var sr []entity.PostList
	for rows.Next() {
		var art entity.PostList
		var cover entity.Cover
		var srCat entity.SearchResultCategory
		var srPt entity.SearchResultPostType
		var publishAt string
		var rel *string
		var lessRel *string
		var username *string
		err := rows.Scan(
			&art.ID,
			&art.Title,
			&art.Slug,
			&publishAt,
			&art.IsCSC,
			&art.PostTypeID,
			&art.CategoryID,
			&art.CoverMediaID,
			&art.CreatorID,
			&cover.UrlMedia,
			&srCat.Name,
			&srCat.Slug,
			&srPt.Name,
			&srPt.Slug,
			&username,
			&rel,
			&lessRel,
		)

		if err != nil {
			panic(err)
		}

		srCat.Url = "/" + srCat.Slug
		art.Category = &srCat
		art.Cover = cover.GetPredefinedSize()
		art.PostType = &srPt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		art.ReleaseDate = t.Unix()

		art.ArticleUrl = "/" + srCat.Slug + "/" + *username + "/" + art.Slug

		sr = append(sr, art)
	}

	return sr, nil
}

/*
	search relevancy ranks:
	1. contains all words from the keyword input (used in select for ordering the result by relevance)
		operators: '+(*word1*) +(*word2*)'
	2. contains at least one word from the keyword input (used in where clause to get all "relevant" data and in select to be the second column to sort by)
		operators: '+(*word1* *word2*)'
*/
func formatKeyword(keyword string) (relevant, lessRelevant string) {
	if len(keyword) == 0 {
		return
	}

	for _, v := range strings.Split(keyword, " ") {
		relevant += "+(*" + v + "*) "
		lessRelevant += "*" + v + "* "
	}

	relevant = "'" + relevant + "'"
	lessRelevant = "'+(" + lessRelevant + ")'"

	return relevant, lessRelevant
}
