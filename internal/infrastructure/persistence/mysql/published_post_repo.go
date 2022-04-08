package mysql

import (
	"strings"
	"time"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PublishedPostRepo struct {
	DB *sqlx.DB
}

func (repo *PublishedPostRepo) GetPublishedPostCountByCategoryId(catId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts
		WHERE category_id = ?
	`

	err := repo.DB.Get(&postCount, query, catId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (repo *PublishedPostRepo) GetPublishedPostCountByTagId(tagId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts pp
		INNER JOIN post_tags pt ON pt.post_id = pp.id 
		WHERE pt.tag_id = ?
	`

	err := repo.DB.Get(&postCount, query, tagId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (repo *PublishedPostRepo) GetPublishedPostCountByPostTypeId(postTypeId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM published_posts pp
		WHERE pp.post_type_id = ?
	`

	err := repo.DB.Get(&postCount, query, postTypeId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (repo *PublishedPostRepo) GetLatestPublishedPost(limit, skip int) ([]entity.PostList, error) {
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

	rows, err := repo.DB.Query(query, limit, skip)
	if err != nil {
		return nil, err
	}

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
}

func (repo *PublishedPostRepo) GetPopularPosts() ([]entity.PostList, error) {
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

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&p.Excerpt,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
}

func (repo *PublishedPostRepo) GetLatestPublishedPostByCategoryId(limit, skip int, categoryId int64) ([]entity.PostList, error) {
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
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
		WHERE pp.category_id = ?
		ORDER BY pp.publish_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := repo.DB.Query(query, categoryId, limit, skip)
	if err != nil {
		return nil, err
	}

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&p.Excerpt,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
}

func (repo *PublishedPostRepo) GetLatestPublishedPostByTagId(limit, skip int, tagId int64) ([]entity.PostList, error) {
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
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
			INNER JOIN post_tags ptg ON ptg.post_id = pp.id
			INNER JOIN tags t ON t.id = ptg.tag_id
		WHERE t.id = ?
		GROUP BY pp.id
		ORDER BY pp.publish_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := repo.DB.Query(query, tagId, limit, skip)
	if err != nil {
		return nil, err
	}

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&p.Excerpt,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
}

func (repo *PublishedPostRepo) GetLatestPublishedPostByCategoryIdAndTagId(limit, skip int, categoryId, tagId int64) ([]entity.PostList, error) {
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
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
			INNER JOIN post_tags ptg ON ptg.post_id = pp.id
			INNER JOIN tags t ON t.id = ptg.tag_id
		WHERE pp.category_id = ?
			AND t.id = ?
		ORDER BY pp.publish_at DESC
		LIMIT ? OFFSET ?
	`

	rows, err := repo.DB.Query(query, categoryId, tagId, limit, skip)
	if err != nil {
		return nil, err
	}

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&p.Excerpt,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
}

func (repo *PublishedPostRepo) SearchPublishedPostByKeyword(keyword string, limit, skip int) ([]entity.PostList, error) {
	relevant, lessRelevant := formatKeyword(keyword)
	stmt, err := repo.DB.Preparex(`
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

	var pl []entity.PostList
	for rows.Next() {
		var p entity.PostList
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var publishAt string
		var rel *string
		var lessRel *string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&cover.UrlMedia,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
			&rel,
			&lessRel,
		)

		if err != nil {
			panic(err)
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat
		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug

		pl = append(pl, p)
	}

	return pl, nil
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

func (repo *PublishedPostRepo) GetPublishedPostDetail(categorySlug, authorUsername, postSlug string) (*entity.PublishedPost, error) {
	query := `
		SELECT
			pp.id, 
			pp.title,
			pp.description,
			pp.slug,
			pp.publish_at,
			pp.is_csc,
			pp.post_type_id,
			pp.category_id,
			pp.cover_media_id,
			pp.created_by,
			pp.excerpt,
			pp.meta_title,
			pp.meta_description,
			m.url_media,
			m.source_name,
			m.source_url,
			m.description,
			m.width,
			m.height,
			m.url_embed,
			c.name,
			c.slug,
			pt.name,
			pt.slug,
			u.username,
			cp.name,
			cp.slug
		FROM published_posts pp
			INNER JOIN post_types pt ON pp.post_type_id = pt.id
			INNER JOIN categories c ON pp.category_id = c.id
			INNER JOIN medias m ON pp.cover_media_id = m.id
			INNER JOIN users u ON pp.created_by = u.id
			LEFT JOIN campaigns cp ON cp.id = pp.campaign_id
		WHERE c.slug = ?
			AND u.username = ?
			AND pp.slug = ?
		LIMIT 1
	`

	rows, err := repo.DB.Query(query, categorySlug, authorUsername, postSlug)
	if err != nil {
		return nil, err
	}

	var p entity.PublishedPost
	for rows.Next() {
		var cover entity.Cover
		var cat entity.CategoryList
		var pt entity.PostTypeList
		var campaign entity.Campaign
		var publishAt string
		var username *string
		err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Description,
			&p.Slug,
			&publishAt,
			&p.IsCSC,
			&p.PostTypeID,
			&p.CategoryID,
			&p.CoverMediaID,
			&p.CreatorID,
			&p.Excerpt,
			&p.MetaTitle,
			&p.MetaDescription,
			&cover.UrlMedia,
			&cover.SourceName,
			&cover.SourceUrl,
			&cover.Description,
			&cover.Width,
			&cover.Height,
			&cover.EmbedVideo,
			&cat.Name,
			&cat.Slug,
			&pt.Name,
			&pt.Slug,
			&username,
			&campaign.Name,
			&campaign.Slug,
		)

		if err != nil {
			panic(err.Error())
		}

		cat.Url = "/" + cat.Slug
		p.Category = &cat

		p.Cover = cover.GetPredefinedSize()
		p.PostType = &pt
		p.Campaign = &campaign

		t, err := time.Parse(time.RFC3339, publishAt)
		if err != nil {
			panic(err)
		}
		p.ReleaseDate = t.Unix()

		p.ArticleUrl = "/" + cat.Slug + "/" + *username + "/" + p.Slug
	}

	return &p, nil
}

func (repo *PublishedPostRepo) IncrementVisitCount(postId int64, updatedAt *time.Time) error {
	tx := repo.DB.MustBegin()
	tx.MustExec("UPDATE published_posts SET visited_count = visited_count + 1, updated_at = ? WHERE id = ?", updatedAt.Format("2006-01-02 15:04:05"), postId)
	return tx.Commit()
}
