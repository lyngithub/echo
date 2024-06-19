package models

import "time"

type Ver struct {
	ID      int64     `xorm:"'id' pk autoincr"` // 主键id
	Created time.Time `xorm:"created"`          // 创建时间
	Updated time.Time `xorm:"updated"`          // 更新时间
	Deleted time.Time `xorm:"deleted"`          // 删除时间
	Key     string    `xorm:"'key'"`            // 当前版本
}

func (v *Ver) TableName() string {
	return "a_ver"
}
