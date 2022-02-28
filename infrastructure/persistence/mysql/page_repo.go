package mysql

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PageRepo struct {
	DB *sqlx.DB
}

func (pageRepo *PageRepo) GetPageBySlug(slug string) (*entity.Page, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.slug,
			p.excerpt,
			p.url,
			p.description,
			p.meta_title,
			p.meta_description
		FROM pages p
		WHERE p.general_status_id = ?
			AND p.slug = ?
		LIMIT 1
	`

	rows, err := pageRepo.DB.Query(query, 1, slug)
	if err != nil {
		return nil, err
	}

	var page = &entity.Page{}
	for rows.Next() {
		err := rows.Scan(
			&page.ID,
			&page.Name,
			&page.Slug,
			&page.Excerpt,
			&page.Url,
			&page.Description,
			&page.MetaTitle,
			&page.MetaDescription,
		)

		if err != nil {
			log.Println("Error scan", err.Error())
		}
	}

	return page, nil
}
