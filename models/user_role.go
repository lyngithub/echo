package models

type UserRole struct {
	UserId int64 `xorm:"'user_id'"` // 用户ID
	RoleId int64 `xorm:"'role_id'"` // 角色ID
}

func (ur *UserRole) TableName() string {
	return "a_user_role"
}
