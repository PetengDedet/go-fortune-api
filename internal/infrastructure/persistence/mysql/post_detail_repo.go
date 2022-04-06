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
			pd.order_num,
			m.url_media,
			m.source_name,
			m.source_url,
			m.description,
			m.width,
			m.height,
			m.url_embed
		FROM post_details pd
		LEFT JOIN medias m ON m.id = pd.content AND (pd.type = 'cover' OR pd.type = 'image')
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
		var cover entity.Cover
		err := rows.Scan(
			&pd.Value,
			&pd.Type,
			&pd.OrderNum,
			&cover.UrlMedia,
			&cover.SourceName,
			&cover.SourceUrl,
			&cover.Description,
			&cover.Width,
			&cover.Height,
			&cover.EmbedVideo,
		)
		if err != nil {
			panic(err)
		}

		if (cover.UrlMedia.String) != "" {
			cover = *cover.GetPredefinedSize()
			pd.Cover = &cover
		}

		postDetails = append(postDetails, pd)
	}

	return postDetails, nil
}
