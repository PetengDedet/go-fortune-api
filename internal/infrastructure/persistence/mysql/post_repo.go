package mysql

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type PostRepo struct {
	DB *sqlx.DB
}

func (repo *PostRepo) GetPostCountByCategoryId(catId int64) (postCount int64, error error) {
	query := `
		SELECT
			COUNT(*) post_count
		FROM posts
		WHERE category_id = ?
	`

	err := repo.DB.Get(&postCount, query, catId)
	if err != nil {
		return 0, err
	}

	return postCount, nil
}

func (repo *PostRepo) IncrementVisitCount(postId int64, updatedAt *time.Time) error {
	tx := repo.DB.MustBegin()
	tx.MustExec("UPDATE posts SET visited_count = visited_count + 1, updated_at = ? WHERE id = ?", updatedAt.Format("2006-01-02 15:04:05"), postId)
	return tx.Commit()
}
