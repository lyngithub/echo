package models

type Role struct {
	Bean        `xorm:"extends"`
	ParentId    int64  `xorm:"'parent_id'"`           // 上级ID
	SuperiorIds string `xorm:"'superior_ids'"`        // 所有上级用户id
	RoleName    string `xorm:"'role_name'"`           // 角色名称
	RoleSort    int    `xorm:"'role_sort'"`           // 显示顺序
	Status      string `xorm:"default('0') 'status'"` // 角色状态（0正常 1停用）
	Remark      string `xorm:"'remark'"`              // 备注
}

func (r *Role) TableName() string {
	return "a_role"
}
