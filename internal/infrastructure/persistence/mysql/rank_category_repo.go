package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type RankCategoryRepo struct {
	DB *sqlx.DB
}

func (rankCategoryRepo *RankCategoryRepo) GetRankCategoryByIds(ids []int64) ([]entity.RankCategory, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			id,
			name,
			slug,
			excerpt,
			is_active,
			meta_title,
			meta_description
		FROM rank_categories
		WHERE id IN(?)
	`, ids)

	if err != nil {
		return nil, err
	}

	query = rankCategoryRepo.DB.Rebind(query)
	rows, err := rankCategoryRepo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var rankCategories []entity.RankCategory
	for rows.Next() {
		var rc entity.RankCategory
		err := rows.Scan(
			&rc.ID,
			&rc.Name,
			&rc.Slug,
			&rc.Excerpt,
			&rc.IsActive,
			&rc.MetaTitle,
			&rc.MetaDescription,
		)

		if err != nil {
			return nil, err
		}
		rankCategories = append(rankCategories, rc)
	}

	return rankCategories, nil
}
