package models

import "time"

type Adminlog struct {
	ID        int64     `xorm:"'id' pk autoincr"` // 主键id
	CreateBy  int64     `xorm:"'create_by'"`      // 创建者
	Created   time.Time `xorm:"created"`          // 创建时间
	Deleted   time.Time `xorm:"deleted"`          // 删除时间
	Method    string    `xorm:"'method'"`         // 请求方式
	Url       string    `xorm:"'url'"`            // 请求地址
	Params    string    `xorm:"'params'"`         // 请求参数
	Name      string    `xorm:"'name'"`           // 标题
	UserName  string    `xorm:"'user_name'"`      // 用户名
	Ip        string    `xorm:"'ip'"`             // ip地址
	UserAgent string    `xorm:"'user_agent'"`     // 浏览器类型和操作系统
}

func (a *Adminlog) TableName() string {
	return "a_adminlog"
}
