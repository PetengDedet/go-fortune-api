package mysql

import (
	"log"

	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/jmoiron/sqlx"
)

type SectionRepo struct {
	DB *sqlx.DB
}

func (sectionRepo *SectionRepo) GetSectionsByPageId(pageId int64) ([]entity.Section, error) {
	query := `
		SELECT 
			s.id,
			s.table_name,
			s.table_id,
			s.title,
			s.slug,
			s.type,
			s.media_id,
			ps.id,
			ps.page_id,
			ps.section_id,
			ps.order_num,
			m.url_media,
			c.name AS category_name,
			c.slug AS category_slug,
			t.name AS tag_name,
			t.slug AS tag_slug,
			pt.name AS post_type_name,
			pt.slug AS post_type_slug,
			lo.url AS linkout_url,
			lo.type AS linkout_type,
			r.name AS rank_name,
			r.slug AS rank_slug,
			rc.name AS rank_category_name,
			rc.slug AS rank_category_slug,
			u.username AS user_username
		FROM sections s
		INNER JOIN page_sections ps ON s.id = ps.section_id
		LEFT JOIN medias m ON s.media_id = m.id
		LEFT JOIN categories c ON s.table_id = c.id AND s.table_name = 'categories'
		LEFT JOIN tags t ON s.table_id = t.id AND s.table_name = 'tags'
		LEFT JOIN post_types pt ON s.table_id = pt.id AND s.table_name = 'post_types'
		LEFT JOIN linkouts lo ON s.table_id = lo.id AND s.table_name = 'linkouts'
		LEFT JOIN ranks r ON s.table_id = r.id AND s.table_name = 'ranks'
		LEFT JOIN rank_categories rc ON s.table_id = rc.id AND s.table_name = 'rank_categories'
		LEFT JOIN users u ON s.table_id = u.id AND s.table_name = 'users'
		WHERE ps.page_id = ?
		ORDER BY ps.order_num ASC
	`

	rows, err := sectionRepo.DB.Query(query, pageId)
	if err != nil {
		return nil, err
	}

	var sections []entity.Section
	for rows.Next() {
		s := &entity.Section{}
		err := rows.Scan(
			&s.ID,
			&s.TableName,
			&s.TableID,
			&s.Title,
			&s.Slug,
			&s.Type,
			&s.MediaID,
			&s.PageSectionID,
			&s.PageSectionPageID,
			&s.PageSectionSectionID,
			&s.OrderNum,
			&s.ImageUrl,
			&s.CategoryName,
			&s.CategorySlug,
			&s.TagName,
			&s.TagSlug,
			&s.PostTypeName,
			&s.PostTypeSlug,
			&s.LinkoutUrl,
			&s.LinkoutType,
			&s.RankName,
			&s.RankSlug,
			&s.RankCategoryName,
			&s.RankCategorySlug,
			&s.UserUsername,
		)

		if err != nil {
			log.Println("Error scan", err.Error())
		}

		sections = append(sections, *s)
	}

	return sections, nil
}

func (sectionRepo *SectionRepo) GetSectionsByPageSlug(slug string) ([]entity.Section, error) {
	query := `
		SELECT 
			s.id,
			s.table_name,
			s.table_id,
			s.title,
			s.slug,
			s.type,
			s.media_id,
			ps.id,
			ps.page_id,
			ps.section_id,
			ps.order_num,
			m.url_media,
			c.name AS category_name,
			c.slug AS category_slug,
			t.name AS tag_name,
			t.slug AS tag_slug,
			pt.name AS post_type_name,
			pt.slug AS post_type_slug,
			lo.url AS linkout_url,
			lo.type AS linkout_type,
			r.name AS rank_name,
			r.slug AS rank_slug,
			rc.name AS rank_category_name,
			rc.slug AS rank_category_slug,
			u.username AS user_username
		FROM sections s
		INNER JOIN page_sections ps ON s.id = ps.section_id
		INNER JOIN pages p ON ps.page_id = p.id
		LEFT JOIN medias m ON s.media_id = m.id
		LEFT JOIN categories c ON s.table_id = c.id AND s.table_name = 'categories'
		LEFT JOIN tags t ON s.table_id = t.id AND s.table_name = 'tags'
		LEFT JOIN post_types pt ON s.table_id = pt.id AND s.table_name = 'post_types'
		LEFT JOIN linkouts lo ON s.table_id = lo.id AND s.table_name = 'linkouts'
		LEFT JOIN ranks r ON s.table_id = r.id AND s.table_name = 'ranks'
		LEFT JOIN rank_categories rc ON s.table_id = rc.id AND s.table_name = 'rank_categories'
		LEFT JOIN users u ON s.table_id = u.id AND s.table_name = 'users'
		WHERE p.slug = ?
		ORDER BY ps.order_num ASC
	`

	rows, err := sectionRepo.DB.Query(query, slug)
	if err != nil {
		return nil, err
	}

	var sections []entity.Section
	for rows.Next() {
		s := &entity.Section{}
		err := rows.Scan(
			&s.ID,
			&s.TableName,
			&s.TableID,
			&s.Title,
			&s.Slug,
			&s.Type,
			&s.MediaID,
			&s.PageSectionID,
			&s.PageSectionPageID,
			&s.PageSectionSectionID,
			&s.OrderNum,
			&s.ImageUrl,
			&s.CategoryName,
			&s.CategorySlug,
			&s.TagName,
			&s.TagSlug,
			&s.PostTypeName,
			&s.PostTypeSlug,
			&s.LinkoutUrl,
			&s.LinkoutType,
			&s.RankName,
			&s.RankSlug,
			&s.RankCategoryName,
			&s.RankCategorySlug,
			&s.UserUsername,
		)

		if err != nil {
			log.Println("Error scan", err.Error())
		}

		sections = append(sections, *s)
	}

	return sections, nil
}
