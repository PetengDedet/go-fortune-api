package mysql

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type CategoryRepo struct {
	DB *sqlx.DB
}

func (categoryRepo *CategoryRepo) GetCategoryBySlug(slug string) (*entity.Category, error) {
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

	var category = &entity.Category{}
	for rows.Next() {
		err := rows.Scan(
			&category.ID,
			&category.Name,
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

func (categoryRepo *CategoryRepo) GetCategoriesByIds(categoryIds []int64) ([]entity.Category, error) {
	if len(categoryIds) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			id,
			name,
			slug,
			excerpt,
			meta_title,
			meta_description
		FROM categories
		WHERE id IN(?)
	`, categoryIds)

	if err != nil {
		return nil, err
	}

	query = categoryRepo.DB.Rebind(query)
	rows, err := categoryRepo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var categories []entity.Category
	for rows.Next() {
		var category entity.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Slug,
			&category.Excerpt,
			&category.MetaTitle,
			&category.MetaDescription,
		)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}
