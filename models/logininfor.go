package models

import "time"

type Logininfor struct {
	ID            int64     `xorm:"'id' pk autoincr"`      // 主键id
	LoginName     string    `xorm:"'login_name'"`          // 登录账号
	Ipaddr        string    `xorm:"'ipaddr'"`              // 登录IP地址
	LoginLocation string    `xorm:"'login_location'"`      // 登录地点
	Browser       string    `xorm:"'browser'"`             // 浏览器类型
	Os            string    `xorm:"'os'"`                  // 操作系统
	Status        string    `xorm:"default('0') 'status'"` // 登录状态（0成功 1失败）
	Msg           string    `xorm:"'msg'"`                 // 提示消息
	Created       time.Time `xorm:"created"`               // 创建时间
}

func (l *Logininfor) TableName() string {
	return "a_logininfor"
}
