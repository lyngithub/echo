package models

type Menu struct {
	Bean           `xorm:"extends"`
	ParentId       int64  `xorm:"'parent_id'"`      // 上级ID
	MenuName       string `xorm:"'menu_name'"`      // 菜单名称
	Weights        int    `xorm:"'weights'"`        // 权重
	Method         string `xorm:"'method'"`         // 请求方式
	Url            string `xorm:"'url'"`            // 请求地址
	Pages          string `xorm:"'pages'"`          // 页面地址
	MenuType       string `xorm:"'menu_type'"`      // 菜单类型（M目录 C菜单 F按钮）
	Classification string `xorm:"'classification'"` // 菜单分类（S查看 E更新 D删除）
	Visible        string `xorm:"'visible'"`        // 菜单状态（0显示 1隐藏）
	Icon           string `xorm:"'icon'"`           // 菜单图标
	Remark         string `xorm:"'remark'"`         // 备注
}

func (r *Menu) TableName() string {
	return "a_menu"
}
