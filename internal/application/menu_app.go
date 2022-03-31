package application

import (
	"strings"

	"github.com/PetengDedet/fortune-post-api/internal/domain/entity"
	"github.com/PetengDedet/fortune-post-api/internal/domain/repository"
)

type MenuApp struct {
	MenuRepo     repository.MenuRepository
	CategoryRepo repository.CategoryRepository
	PageRepo     repository.PageRepository
	LinkoutRepo  repository.LinkoutRepository
	RankRepo     repository.RankRepository
}

func (app *MenuApp) GetPublicMenuPositions() ([]entity.MenuPosition, error) {
	var menuPositions []entity.MenuPosition

	menuPost, err := app.MenuRepo.GetMenuPositions()
	if err != nil {
		return nil, err
	}

	if len(menuPost) == 0 {
		return []entity.MenuPosition{}, nil
	}

	positionIds := getMenuPositionIds(menuPost)

	parentMenus, err := app.MenuRepo.GetMenusByPositionIds(positionIds)
	if err != nil {
		return nil, err
	}
	//TODO: if no parent menus, return menupositions without menu

	catIds, loIds, pageIds, rankIds := getMenuRelationIds(parentMenus)
	parentMenus = mapMenuType(
		app,
		parentMenus,
		catIds,
		loIds,
		pageIds,
		rankIds,
	)

	parentMenuIds := getParentMenuIds(parentMenus)
	childrenMenus, err := app.MenuRepo.GetChildrenMenus(parentMenuIds)
	if err != nil {
		return nil, err
	}

	cIds, lIds, pIds, rkIds := getMenuRelationIds(childrenMenus)
	childrenMenus = mapMenuType(
		app,
		childrenMenus,
		cIds,
		lIds,
		pIds,
		rkIds,
	)

	parentMenus = mapChildMenusToParentMenus(parentMenus, childrenMenus)

	for _, mp := range menuPost {
		mp.GetMenus(parentMenus)
		mp.Position = strings.ToLower(mp.Position)
		menuPositions = append(menuPositions, mp)
	}

	return menuPositions, nil
}

func getMenuPositionIds(mp []entity.MenuPosition) (positionIds []int) {
	for _, mp := range mp {
		positionIds = append(positionIds, int(mp.ID))
	}

	return positionIds
}

func getParentMenuIds(pm []entity.Menu) (parentMenuIds []int) {
	for _, pm := range pm {
		parentMenuIds = append(parentMenuIds, int(pm.ID))
	}

	return parentMenuIds
}

func getMenuRelationIds(m []entity.Menu) (catIds, loIds, pageIds, rankIds []int64) {
	//TODO: contine on each match
	for _, m := range m {
		if m.TableName.String == "categories" && m.TableID.Int64 != 0 {
			catIds = append(catIds, m.TableID.Int64)
		}

		if m.TableName.String == "linkouts" && m.TableID.Int64 != 0 {
			loIds = append(loIds, m.TableID.Int64)
		}

		if m.TableName.String == "pages" && m.TableID.Int64 != 0 {
			pageIds = append(pageIds, m.TableID.Int64)
		}

		if m.TableName.String == "ranks" && m.TableID.Int64 != 0 {
			rankIds = append(rankIds, m.TableID.Int64)
		}
	}

	return catIds, loIds, pageIds, rankIds
}

func mapMenuType(app *MenuApp, menus []entity.Menu, catIds, loIds, pageIds, rankIds []int64) []entity.Menu {
	categories, err := app.CategoryRepo.GetCategoriesByIds(catIds)
	if err != nil {
		panic(err.Error())
	}

	linkouts, err := app.LinkoutRepo.GetLinkoutsByIds(loIds)
	if err != nil {
		panic(err.Error())
	}

	pages, err := app.PageRepo.GetPagesByIds(pageIds)
	if err != nil {
		panic(err.Error())
	}

	ranks, err := app.RankRepo.GetRanksByIds(rankIds)
	if err != nil {
		panic(err.Error())
	}

	for index, m := range menus {
		if m.TableName.String == "categories" && m.TableID.Int64 != 0 {
			for _, cat := range categories {
				if cat.ID.Int64 == m.TableID.Int64 {
					m.Category = &cat
					menus[index] = *m.SetMenuAttributes()
					break
				}
			}
			continue
		}

		if m.TableName.String == "linkouts" && m.TableID.Int64 != 0 {
			for _, lo := range linkouts {
				if lo.ID == m.TableID.Int64 {
					m.Linkout = &lo
					menus[index] = *m.SetMenuAttributes()
					break
				}
			}
			continue
		}

		if m.TableName.String == "pages" && m.TableID.Int64 != 0 {
			for _, pg := range pages {
				if pg.ID == m.TableID.Int64 {
					m.Page = &pg
					menus[index] = *m.SetMenuAttributes()
					break
				}
			}
			continue
		}

		if m.TableName.String == "ranks" && m.TableID.Int64 != 0 {
			for _, rk := range ranks {
				if rk.ID == m.TableID.Int64 {
					m.Rank = &rk
					menus[index] = *m.SetMenuAttributes()
					break
				}
			}
			continue
		}

		menus[index] = *m.SetMenuAttributes()
	}

	return menus
}

func mapChildMenusToParentMenus(parent, childs []entity.Menu) []entity.Menu {
	for index, p := range parent {
		parent[index].ChildMenu = p.GetChildMenus(childs)
	}

	return parent
}
