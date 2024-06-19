package models

import "time"

type User struct {
	Bean        `xorm:"extends"`
	ParentId    int64     `xorm:"'parent_id'"`               // 上级ID
	SuperiorIds string    `xorm:"'superior_ids'"`            // 所有上级用户id
	LoginName   string    `xorm:"'login_name'"`              // 登录账号
	GoogleCode  string    `xorm:"'google_code'"`             // 谷歌验证码
	Username    string    `xorm:"'user_name'"`               // 用户昵称
	UserType    string    `xorm:"default('01') 'user_type'"` // 用户类型（00系统用户 01注册用户）
	Email       string    `xorm:"'email'"`                   // 用户邮箱
	Phonenumber string    `xorm:"'phonenumber'"`             // 手机号码
	Password    string    `xorm:"'password'"`                // 密码
	Status      string    `xorm:"default('0') 'status'"`     // 帐号状态（0正常 1停用）
	LoginIp     string    `xorm:"'login_ip'"`                // 最后登陆IP
	LoginDate   time.Time `xorm:"'login_date'"`              // 最后登陆时间
	Remark      string    `xorm:"'remark'"`                  // 备注
}

func (u *User) TableName() string {
	return "a_user"
}
