package models

type RoleMenu struct {
	RoleId int64 `xorm:"'role_id'"` // 角色ID
	MenuId int64 `xorm:"'menu_id'"` // 菜单ID
}

func (rm *RoleMenu) TableName() string {
	return "a_role_menu"
}
