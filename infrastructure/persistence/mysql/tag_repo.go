package mysql

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type TagRepo struct {
	DB *sqlx.DB
}

func (tagRepo *TagRepo) GetTagByIds(ids []int64) ([]entity.Tag, error) {
	if len(ids) == 0 {
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
		FROM tags
		WHERE id IN(?)
	`, ids)

	if err != nil {
		return nil, err
	}

	query = tagRepo.DB.Rebind(query)
	rows, err := tagRepo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var tags []entity.Tag
	for rows.Next() {
		var t entity.Tag
		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Slug,
			&t.Excerpt,
			&t.MetaTitle,
			&t.MetaDescription,
		)

		if err != nil {
			return nil, err
		}
		tags = append(tags, t)
	}

	return tags, nil
}

func (tagRepo *TagRepo) GetTagBySlug(slug string) (*entity.Tag, error) {
	query := `
		SELECT
			id,
			name,
			slug,
			excerpt,
			meta_title,
			meta_description
		FROM tags
		WHERE slug = ?
		LIMIT 1
	`

	rows, err := tagRepo.DB.Query(query, slug)
	if err != nil {
		return nil, err
	}

	var t = &entity.Tag{}
	for rows.Next() {
		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Slug,
			&t.Excerpt,
			&t.MetaTitle,
			&t.MetaDescription,
		)

		if err != nil {
			return nil, err
		}
	}

	return t, nil
}
