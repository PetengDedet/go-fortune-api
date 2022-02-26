package entity

type PublicMenu struct {
	Name       string       `json:"name"`
	Slug       string       `json:"slug"`
	Type       string       `json:"type"`
	Url        string       `json:"url"`
	OrderNum   int64        `json:"order_num"`
	IsActive   bool         `json:"is_active"`
	ChildMenu  []PublicMenu `json:"child_menu"`
	LinkoutUrl string       `json:"-"`
}

func PublicMenuResponse(pm PublicMenu) PublicMenu {
	pm.Url = "/" + pm.Type
	if pm.Type == "linkouts" {
		pm.Url = pm.LinkoutUrl
		pm.Type = "link"
	}

	if pm.Type == "categories" {
		pm.Type = "category"
		pm.Url = "/" + pm.Slug
	}

	if pm.Type == "pages" {
		pm.Type = "page"
	}

	if pm.Type == "post_types" {
		pm.Type = "content"
	}

	if pm.Type == "tags" {
		pm.Type = "tag"
	}

	if pm.Type == "ranks" {
		pm.Type = "rank"
	}

	if len(pm.ChildMenu) <= 0 || pm.ChildMenu == nil {
		pm.ChildMenu = []PublicMenu{}
	}

	if len(pm.ChildMenu) > 0 {
		pm.Url = ""
	}

	return pm
}

type Menu struct {
	ID             int64
	Title          string
	Slug           string
	ParentMenuID   int64
	OrderNum       int64
	MenuPositionID int64
	MenuType       string
	LinkoutUrl     string
	IsActive       int64
}
