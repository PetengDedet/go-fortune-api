package mysql

import (
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type PostTypeRepo struct {
	DB *sqlx.DB
}

func (postTypeRepo *PostTypeRepo) GetPostTypeByIds(ids []int64) ([]entity.PostType, error) {
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
		FROM post_types
		WHERE id IN(?)
	`, ids)

	if err != nil {
		return nil, err
	}

	query = postTypeRepo.DB.Rebind(query)
	rows, err := postTypeRepo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var postTypes []entity.PostType
	for rows.Next() {
		var pt entity.PostType
		err := rows.Scan(
			&pt.ID,
			&pt.Name,
			&pt.Slug,
			&pt.Excerpt,
			&pt.MetaTitle,
			&pt.MetaDescription,
		)

		if err != nil {
			return nil, err
		}
		postTypes = append(postTypes, pt)
	}

	return postTypes, nil
}
