package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type RankRepo struct {
	DB *sqlx.DB
}

func (repo *RankRepo) GetRanksByIds(rankIds []int64) ([]entity.Rank, error) {
	if len(rankIds) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			id,
			name,
			slug,
			description,
			title
		FROM ranks
		WHERE id IN(?)
	`, rankIds)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var ranks []entity.Rank
	for rows.Next() {
		var rank entity.Rank
		err := rows.Scan(
			&rank.ID,
			&rank.Name,
			&rank.Slug,
			&rank.Description,
			&rank.Title,
		)

		if err != nil {
			return nil, err
		}
		ranks = append(ranks, rank)
	}

	return ranks, nil
}
