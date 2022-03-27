package mysql

import (
	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	DB *sqlx.DB
}

func (pr *PostRepo) GetPostCountByCategoryId(catId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM posts
		WHERE category_id = ?
	`

	err := pr.DB.Get(&postCount, query, catId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}
