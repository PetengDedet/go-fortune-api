package mysql

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PageRepo struct {
	DB *sqlx.DB
}

func (repo *PageRepo) GetPageBySlug(slug string) (*entity.Page, error) {
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

	rows, err := repo.DB.Query(query, 1, slug)
	if err != nil {
		return nil, err
	}

	var page = &entity.Page{}
	for rows.Next() {
		err := rows.Scan(
			&page.ID,
			&page.Page,
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

func (repo *PageRepo) GetPagesByIds(pageIds []int64) ([]entity.Page, error) {
	if len(pageIds) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			id,
			name,
			slug,
			excerpt,
			url,
			description,
			meta_title,
			meta_description
		FROM pages
		WHERE id IN(?)
	`, pageIds)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var pages []entity.Page
	for rows.Next() {
		var page entity.Page
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
			return nil, err
		}
		pages = append(pages, page)
	}

	return pages, nil
}
