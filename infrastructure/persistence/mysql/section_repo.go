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
			m.url_media
		FROM sections s
		INNER JOIN page_sections ps ON s.id = ps.section_id
		LEFT JOIN medias m ON s.media_id = m.id
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
		m := &entity.Media{}
		err := rows.Scan(
			&s.ID,
			&s.TableName,
			&s.TableID,
			&s.Title,
			&s.Slug,
			&s.Type,
			&m.ID,
			&s.PageSectionID,
			&s.PageSectionPageID,
			&s.PageSectionSectionID,
			&s.OrderNum,
			&m.UrlMedia,
		)

		if err != nil {
			log.Println("Error scan", err.Error())
		}

		if m.ID.Int64 != 0 {
			s.Media = m
		}

		sections = append(sections, *s)
	}

	return sections, nil
}
