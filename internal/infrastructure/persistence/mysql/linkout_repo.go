package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type LinkoutRepo struct {
	DB *sqlx.DB
}

func (repo *LinkoutRepo) GetLinkoutsByIds(linkoutIds []int64) ([]entity.Linkout, error) {
	if len(linkoutIds) == 0 {
		return nil, nil
	}

	query, args, err := sqlx.In(`
		SELECT
			id,
			url,
			type
		FROM linkouts
		WHERE id IN(?)
	`, linkoutIds)

	if err != nil {
		return nil, err
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var linkouts []entity.Linkout
	for rows.Next() {
		var linkout entity.Linkout
		err := rows.Scan(
			&linkout.ID,
			&linkout.Url,
			&linkout.Type,
		)
		if err != nil {
			return nil, err
		}
		linkouts = append(linkouts, linkout)
	}

	return linkouts, nil
}

func (repo *LinkoutRepo) GetLinkoutsByType(t string) ([]entity.Linkout, error) {
	query := `
		SELECT
			lo.id,
			lo.url,
			lo.type
		FROM linkouts lo
		WHERE lo.type = ?
	`
	rows, err := repo.DB.Query(query, t)
	if err != nil {
		return nil, err
	}

	var linkouts []entity.Linkout
	for rows.Next() {
		var linkout entity.Linkout
		err := rows.Scan(
			&linkout.ID,
			&linkout.Url,
			&linkout.Type,
		)
		if err != nil {
			return nil, err
		}
		linkouts = append(linkouts, linkout)
	}

	return linkouts, nil
}
