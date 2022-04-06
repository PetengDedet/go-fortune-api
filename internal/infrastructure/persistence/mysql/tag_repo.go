package mysql

import (
	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/jmoiron/sqlx"
	"gopkg.in/guregu/null.v4"
)

type TagRepo struct {
	DB *sqlx.DB
}

func (repo *TagRepo) GetTagByIds(ids []int64) ([]entity.Tag, error) {
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

	query = repo.DB.Rebind(query)
	rows, err := repo.DB.Query(query, args...)
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

func (repo *TagRepo) GetTagBySlug(slug string) (*entity.Tag, error) {
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

	rows, err := repo.DB.Query(query, slug)
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

func (repo *TagRepo) GetPostIdsByTagId(id int64) ([]int64, error) {
	query := "SELECT post_id FROM post_tags WHERE tag_id = ?"
	rows, err := repo.DB.Query(query, id)
	if err != nil {
		return nil, err
	}

	var postIds = []int64{}
	for rows.Next() {
		var pid int64
		err := rows.Scan(&pid)
		if err != nil {
			panic(err)
		}
		postIds = append(postIds, pid)
	}

	return postIds, nil
}

func (repo *TagRepo) GetTagsByPostId(postId int64) ([]entity.Tag, error) {
	query := `
		SELECT
			t.name,
			t.slug,
			pt.order_num
		FROM post_tags pt
		INNER JOIN tags t ON pt.tag_id = t.id
		WHERE pt.post_id = ?
		ORDER BY pt.order_num
	`
	rows, err := repo.DB.Query(query, postId)
	if err != nil {
		return nil, err
	}

	var tags []entity.Tag
	for rows.Next() {
		var t entity.Tag
		err := rows.Scan(&t.Name, &t.Slug, &t.OrderNum)
		if err != nil {
			panic(err)
		}

		t.Url = null.StringFrom("/tag/" + t.Slug)

		tags = append(tags, t)
	}

	return tags, nil
}
