package mysql

import (
	"strconv"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

type MagazineRepo struct {
	DB *sqlx.DB
}

func (repo *MagazineRepo) GetLatestActiveMagazines(limit int) ([]entity.Magazine, error) {
	months := []string{
		"January", "February", "March",
		"April", "May", "June",
		"July", "August", "September",
		"October", "November", "December",
	}

	query := `
		SELECT
			mz.id,
			mz.title,
			mz.slug,
			mz.description,
			mz.edition_month,
			mz.edition_year,
			mz.buy_url,
			m.url_media,
			m.source_url,
			m.source_name,
			m.description,
			m.url_embed
		FROM magazines mz
		INNER JOIN medias m ON m.id = mz.cover
		WHERE is_active = 1
		ORDER BY mz.edition_year DESC, edition_month DESC
		LIMIT ?
	`

	rows, err := repo.DB.Query(query, limit)
	if err != nil {
		return nil, err
	}

	var magazines []entity.Magazine
	for rows.Next() {
		var mz entity.Magazine
		var m entity.Media
		var year int
		var month int

		err := rows.Scan(
			&mz.ID,
			&mz.Title,
			&mz.Slug,
			&mz.Description,
			&month,
			&year,
			&mz.PurchaseLink,
			&m.UrlMedia,
			&m.SourceUrl,
			&m.SourceName,
			&m.Description,
			&m.UrlEmbed,
		)

		if err != nil {
			panic(err)
		}

		mz.Edition = null.StringFrom(months[month-1] + " " + strconv.Itoa(year))
		m = *m.SetUrl("600", "auto")
		mz.Cover = &entity.MagazineCover{
			ImageUrl:    m.Url,
			SourceName:  m.SourceName,
			SourceUrl:   m.SourceUrl,
			Description: m.Description,
			EmbedVideo:  m.UrlEmbed,
		}
		magazines = append(magazines, mz)
	}

	return magazines, nil
}
