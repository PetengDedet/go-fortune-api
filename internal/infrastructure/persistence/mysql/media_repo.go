package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type MediaRepo struct {
	DB *sqlx.DB
}

func (repo *MediaRepo) GetMediaByIds(ids []int64) ([]entity.Media, error) {
	query, args, err := sqlx.In(`
		SELECT
			m.id,
			m.gallery_id,
			m.name,
			m.description,
			m.url_media,
			m.mime,
			m.extension,
			m.source_name,
			m.source_url,
			m.width,
			m.height,
			m.keyword,
			m.url_embed
		FROM medias m
		INNER JOIN galleries g ON m.gallery_id = g.id
		WHERE m.id IN (?)
	`, ids)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var medias []entity.Media
	for rows.Next() {
		var m entity.Media
		err := rows.Scan(
			&m.ID,
			&m.GalleryID,
			&m.Name,
			&m.Description,
			&m.UrlMedia,
			&m.Mime,
			&m.Extension,
			&m.SourceName,
			&m.SourceUrl,
			&m.Width,
			&m.Height,
			&m.Keyword,
			&m.UrlEmbed,
		)
		if err != nil {
			panic(err)
		}

		medias = append(medias, m)
	}

	return medias, nil
}
