package mysql

import (
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
