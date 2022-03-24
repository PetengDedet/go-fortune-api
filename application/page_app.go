package application

import (
	"net/url"
	"os"
	"strconv"

	"github.com/PetengDedet/fortune-post-api/common"
	"github.com/PetengDedet/fortune-post-api/domain/entity"
	"github.com/PetengDedet/fortune-post-api/domain/repository"
	"gopkg.in/guregu/null.v4"
)

type PageAppInterface interface {
	GetPageDetailBySlug(slug string) (*entity.Page, error)
}
type PageApp struct {
	PageRepo         repository.PageRepository
	SectionRepo      repository.SectionRepository
	CategoryRepo     repository.CategoryRepository
	LinkoutRepo      repository.LinkoutRepository
	PostTypeRepo     repository.PostTypeRepository
	TagRepo          repository.TagRepository
	RankRepo         repository.RankRepository
	RankCategoryRepo repository.RankCategoryRepository
}

// var _ PageAppInterface = &PageApp{}

func (pageApp *PageApp) GetPageDetailBySlug(slug string) (*entity.Page, error) {
	page, err := pageApp.PageRepo.GetPageBySlug(slug)
	if err != nil {
		return nil, err
	}

	// Empty page
	if page.ID == 0 {
		return nil, &common.NotFoundError{}
	}

	sections, err := pageApp.SectionRepo.GetSectionsByPageId(page.ID)
	if err != nil {
		return nil, err
	}

	if len(sections) == 0 {
		page.Sections = []entity.Section{}
		return page, nil
	}

	catIds, loIds, ptIds, tIds, rIds, rcIds := getSectionRelationIds(sections)
	sections = mapSectionType(
		pageApp,
		sections,
		catIds,
		loIds,
		ptIds,
		tIds,
		rIds,
		rcIds,
	)

	page.Sections = sections

	return page, nil
}

func (pageApp *PageApp) GetCategoryPageDetail(slug string, category *entity.Category) (*entity.Page, error) {
	page, err := pageApp.GetPageDetailBySlug(slug)
	if err != nil {
		return nil, err
	}

	page.Page = null.StringFrom(category.Name)
	page.Slug = null.StringFrom(category.Slug)
	page.Excerpt = category.Excerpt
	page.MetaTitle = category.MetaTitle
	page.MetaDescription = category.MetaDescription
	page.Url = null.StringFrom("/" + category.Slug)
	page.ArticleCounts = category.PublishedPostCount

	// Set section url
	for i, s := range page.Sections {
		if s.Type.String == "headline" && s.Slug.String == "headline-categories" {
			page.Sections[i] = *s.MutateUrl("/v1/headline?category=" + category.Slug)
			page.Sections[i] = *s.MutateBaseUrl("/" + category.Slug)
			continue
		}

		if s.Type.String == "latest" && s.Slug.String == "latest-categories" {
			page.Sections[i] = *s.MutateUrl("/v1/latest/category/" + category.Slug)
			page.Sections[i] = *s.MutateBaseUrl("/" + category.Slug)
			continue
		}
	}

	return page, nil
}

func (pageApp *PageApp) GetTagPageDetail(slug string, tag *entity.Tag) (*entity.Page, error) {
	page, err := pageApp.GetPageDetailBySlug(slug)
	if err != nil {
		return nil, err
	}

	page.Page = null.StringFrom(tag.Name)
	page.Slug = null.StringFrom(tag.Slug)
	page.Excerpt = tag.Excerpt
	page.MetaTitle = tag.MetaTitle
	page.MetaDescription = tag.MetaDescription
	page.Url = null.StringFrom("/tag/" + tag.Slug)
	page.ArticleCounts = tag.PublishedPostCount

	// Set section url
	for i, s := range page.Sections {
		if s.Type.String == "latest" && s.Slug.String == "latest-tags" {
			page.Sections[i] = *s.MutateUrl("/v1/latest/tag/" + tag.Slug)
			page.Sections[i] = *s.MutateBaseUrl("/tag/" + tag.Slug)
			continue
		}
	}

	return page, nil
}

func (pageApp *PageApp) GetSearchResultPageDetail(keyword string, currentPage int, searchResult []entity.SearchResultArticle) (*entity.Page, error) {
	page, err := pageApp.GetPageDetailBySlug("search")
	if err != nil {
		return nil, err
	}

	domain := os.Getenv("APP_URL")
	if len(domain) == 0 {
		domain = "https://post-api.fortuneidn.com/"
	}

	page.MetaTitle = null.StringFrom("Berita " + strconv.Quote(keyword) + " Terkini Hari Ini")
	page.MetaDescription = null.StringFrom("Kumpulan berita " + strconv.Quote(keyword) + " terkini hari ini bisa kamu simak dengan lengkap disini dari berbagai macam sudut pandang.")

	pageQueryParams := url.Values{
		"keyword": {keyword},
		"page":    {strconv.Itoa(currentPage)},
	}

	page.ArticleCounts = int64(len(searchResult))
	page.Url = null.StringFrom(page.Url.String + "?" + pageQueryParams.Encode())

	nextUrlQueryParams := url.Values{
		"keyword": {keyword},
		"page":    {strconv.Itoa(currentPage + 1)},
	}
	page.Articles = &entity.SearchResultArticles{
		NextUrl: null.StringFrom(domain + "/search?" + nextUrlQueryParams.Encode()),
		Data:    searchResult,
	}

	return page, nil
}

func getSectionRelationIds(s []entity.Section) (catIds, loIds, ptIds, tIds, rIds, rcIds []int64) {
	for _, s := range s {
		if s.TableName.String == "categories" && s.TableID.Int64 != 0 {
			catIds = append(catIds, s.TableID.Int64)
			continue
		}

		if s.TableName.String == "linkouts" && s.TableID.Int64 != 0 {
			loIds = append(loIds, s.TableID.Int64)
			continue
		}

		if s.TableName.String == "ranks" && s.TableID.Int64 != 0 {
			rIds = append(rIds, s.TableID.Int64)
			continue
		}

		if s.TableName.String == "post_sections" && s.TableID.Int64 != 0 {
			ptIds = append(ptIds, s.TableID.Int64)
			continue
		}

		if s.TableName.String == "rank_categories" && s.TableID.Int64 != 0 {
			rcIds = append(rcIds, s.TableID.Int64)
			continue
		}

		if s.TableName.String == "tags" && s.TableID.Int64 != 0 {
			tIds = append(tIds, s.TableID.Int64)
			continue
		}
	}

	return catIds, loIds, ptIds, tIds, rIds, rcIds
}

func mapSectionType(pageApp *PageApp, sections []entity.Section, catIds, loIds, ptIds, tIds, rIds, rcIds []int64) []entity.Section {
	categories, err := pageApp.CategoryRepo.GetCategoriesByIds(catIds)
	if err != nil {
		panic(err.Error())
	}

	linkouts, err := pageApp.LinkoutRepo.GetLinkoutsByIds(loIds)
	if err != nil {
		panic(err.Error())
	}

	tags, err := pageApp.TagRepo.GetTagByIds(tIds)
	if err != nil {
		panic(err.Error())
	}

	ranks, err := pageApp.RankRepo.GetRanksByIds(rIds)
	if err != nil {
		panic(err.Error())
	}

	rankCategories, err := pageApp.RankCategoryRepo.GetRankCategoryByIds(rcIds)
	if err != nil {
		panic(err.Error())
	}

	postTypes, err := pageApp.PostTypeRepo.GetPostTypeByIds(ptIds)
	if err != nil {
		panic(err.Error())
	}

	for index, s := range sections {
		if s.TableName.String == "categories" && s.TableID.Int64 != 0 {
			for _, cat := range categories {
				if cat.ID.Int64 == s.TableID.Int64 {
					s.Category = &cat
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		if s.TableName.String == "tags" && s.TableID.Int64 != 0 {
			for _, tag := range tags {
				if tag.ID == s.TableID.Int64 {
					s.Tag = &tag
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		if s.TableName.String == "linkouts" && s.TableID.Int64 != 0 {
			for _, lo := range linkouts {
				if lo.ID == s.TableID.Int64 {
					s.Linkout = &lo
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		if s.TableName.String == "ranks" && s.TableID.Int64 != 0 {
			for _, r := range ranks {
				if r.ID == s.TableID.Int64 {
					s.Rank = &r
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		if s.TableName.String == "rank_categories" && s.TableID.Int64 != 0 {
			for _, rc := range rankCategories {
				if rc.ID == s.TableID.Int64 {
					s.RankCategory = &rc
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		if s.TableName.String == "post_types" && s.TableID.Int64 != 0 {
			for _, pt := range postTypes {
				if pt.ID == s.TableID.Int64 {
					s.PostType = &pt
					sections[index] = *s.SetSectionAttributes()
					break
				}
			}
			continue
		}

		sections[index] = *s.SetSectionAttributes()
	}

	return sections
}
