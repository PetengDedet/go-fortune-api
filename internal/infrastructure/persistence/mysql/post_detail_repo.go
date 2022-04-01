package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PostDetailRepo struct {
	DB *sqlx.DB
}

func (repo *PostDetailRepo) GetPostDetailsByPostId(postId int64) ([]entity.PostDetailList, error) {
	query := `
		SELECT
			pd.content,
			pd.type,
			pd.order_num
		FROM post_details pd
		WHERE pd.post_id = ?
			AND pd.deleted_at IS NULL
		ORDER BY order_num
	`

	rows, err := repo.DB.Query(query, postId)
	if err != nil {
		return nil, err
	}

	var postDetails []entity.PostDetailList
	for rows.Next() {
		var pd entity.PostDetailList
		err := rows.Scan(
			&pd.Value,
			&pd.Type,
			&pd.OrderNum,
		)
		if err != nil {
			panic(err)
		}

		postDetails = append(postDetails, pd)
	}

	return postDetails, nil
}
