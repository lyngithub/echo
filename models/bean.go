package models

import "time"

type Bean struct {
	ID       int64     `xorm:"'id' pk autoincr"`          // 主键id
	CreateBy int64     `xorm:"'create_by'"`               // 创建者
	Created  time.Time `xorm:"created"`                   // 创建时间
	UpdateBy int64     `xorm:"'update_by'"`               // 更新者
	Updated  time.Time `xorm:"updated"`                   // 更新时间
	Deleted  time.Time `xorm:"deleted"`                   // 删除时间
	Version  int64     `xorm:"notnull version 'version'"` // 版本
}

func Add(id int64, bean *Bean) {
	bean.CreateBy = id
	bean.UpdateBy = id
}

func Edit(id int64, bean *Bean) {
	bean.UpdateBy = id
}
