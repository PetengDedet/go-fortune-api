package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type KeywordRepo struct {
	DB *sqlx.DB
}

func (repo *KeywordRepo) SaveNewKeyword(keyword string) error {
	return nil
}

func (repo *KeywordRepo) GetPopularKeyword() ([]entity.Keyword, error) {
	query := "SELECT keyword FROM popular_keywords"
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err
	}

	var kw []entity.Keyword
	for rows.Next() {
		var k *entity.Keyword
		err := rows.Scan(&k)
		if err != nil {
			panic(err)
		}

		kw = append(kw, *k)
	}

	return kw, nil
}
