package models

type IpWhitelist struct {
	Bean      `xorm:"extends"`
	CompanyId int64  `xorm:"'company_id' comment('企业id')"` // 企业id
	Ip        string `xorm:"'ip' comment('ip(多ip用，分割)')"`  // ip(多ip用，分割)
	Remark    string `xorm:"'remark' comment('备注')"`       // 备注
}

func (c *IpWhitelist) TableName() string {
	return "a_ip_whitelist"
}
