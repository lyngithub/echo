package vo

type MenuTreeVo struct {
	Id       int64         `json:"id"`
	Name     string        `json:"name"`
	Weights  int           `json:"weights"`
	Icon     string        `json:"icon"`
	MenuType string        `json:"menu_type"`
	Pages    string        `json:"pages"`
	Children []*MenuTreeVo `json:"children"`
}

type RoleTreeVo struct {
	Id       int64         `json:"id"`
	Name     string        `json:"name"`
	Children []*RoleTreeVo `json:"children"`
}
