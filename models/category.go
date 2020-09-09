package models

import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id     int
	Pid    int    //分类id
	Title  string `orm:"size(30);unique"`
	Intro  string //介绍
	Icon   string
	Cnt    int  //统计分类下图书
	Sort   int  //排序
	Status bool //状态，true 显示
}

func (m *Category) TableName() string {
	return TNCategory()
}

func (m *Category) FindAll(pid int, status int) (cates []Category, err error) {
	qs := orm.NewOrm().QueryTable(m.TableName())
	if pid > -1 {
		qs = qs.Filter("pid", pid)
	}
	if status == 0 || status == 1 {
		qs = qs.Filter("status", status)
	}
	_, err = qs.OrderBy("id").All(&cates)
	return
}

func (m *Category) Find(id int) (cate Category, err error) {
	cate.Id = id
	if err := orm.NewOrm().Read(&cate); err != nil {
		return cate, err
	}
	return
}
