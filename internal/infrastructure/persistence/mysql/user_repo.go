package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	DB *sqlx.DB
}

func (repo *UserRepo) GetAuthorsByPostIds(postIds []int64) ([]entity.Author, error) {
	query, args, err := sqlx.In(`
		SELECT
			u.name,
			u.username,
			u.nickname,
			u.avatar,
			post_id
		FROM post_authors pa
		INNER JOIN users u ON pa.author_id = u.id
		WHERE post_id IN (?)
		ORDER BY order_num
	`, postIds)

	if err != nil {
		panic(err)
	}

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}

	var authors []entity.Author
	for rows.Next() {
		var author entity.Author
		err := rows.Scan(
			&author.Name,
			&author.Username,
			&author.Nickname,
			&author.Avatar,
			&author.PostID,
		)
		if err != nil {
			panic(err)
		}

		author.AuthorUrl = "/" + author.Username

		authors = append(authors, author)
	}

	return authors, nil
}

func (repo *UserRepo) GetAuthorsByPostId(postId int64) (*entity.Author, error) {
	return nil, nil
}
