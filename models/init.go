package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(Member),
	)
}

// 多个数据库怎么获取表前缀
//func TablePrefix() string {
//	return beego.AppConfig.String("")
//}

func TNMembers() string {
	return "md_members"
}
