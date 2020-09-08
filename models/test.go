package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Test struct {
	Id        int       `orm:"pk;auto" json:"id"`
	Title     string    `json:"title"`
	Status    int8      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//func init() {
//	//orm.RegisterDriver("mysql", orm.DRMySQL)
//	dbHost := beego.AppConfig.String("db_w_host")
//	dbPort := beego.AppConfig.String("db_w_port")
//	dbUser := beego.AppConfig.String("db_w_user")
//	dbPwd := beego.AppConfig.String("db_w_password")
//	dbName := beego.AppConfig.String("db_w_database")
//	dbPrefix := beego.AppConfig.String("db_w_prefix")
//	dbCharset := beego.AppConfig.String("db_w_charset")
//	ds := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=" + dbCharset
//	orm.RegisterDataBase("default", "mysql", ds, 10, 10)
//	orm.RegisterModelWithPrefix(dbPrefix, new(Test))
//}

//crud
func (m *Test) GetById(id int) (test Test) {
	//test.Id = id
	//orm.NewOrm().Read(&test)
	//return test

	t := Test{Id: id}
	o := orm.NewOrm()
	err := o.Read(&t)
	if err == nil {
		return t
	} else {
		return
	}
}
