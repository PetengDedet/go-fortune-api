package mysql

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	DB *sqlx.DB
}

func (categoryRepo *CategoryRepo) GetCategoryPageBySlug(slug string) (*entity.CategoryPage, error) {
	query := `
		SELECT 
			c.id,
			c.name,
			c.slug,
			c.excerpt,
			c.meta_title,
			c.meta_description
		FROM categories c
		WHERE c.slug = ?
			AND c.general_status_id = ?
		LIMIT 1
	`

	rows, err := categoryRepo.DB.Query(query, slug, 1)
	if err != nil {
		return nil, err
	}

	var category = &entity.CategoryPage{}
	for rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Page,
			&category.Slug,
			&category.Excerpt,
			&category.MetaTitle,
			&category.MetaDescription,
		)

		if err != nil {
			log.Println("Error scan", err.Error())
		}
	}

	return category, nil
}
