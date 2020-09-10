package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"strings"
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

// 获取所有分类
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

// 根据id获取分类
func (m *Category) Find(id int) (cate Category, err error) {
	cate.Id = id
	if err := orm.NewOrm().Read(&cate); err != nil {
		return cate, err
	}
	return
}

func (m *Category) MultiCreate(pid int, cates string) (err error) {
	slice := strings.Split(cates, "\n")
	if 0 == len(slice) {
		return errors.New("参数异常")
	}

	o := orm.NewOrm()
	for _, item := range slice {
		if item = strings.TrimSpace(item); "" != item {
			var cate = Category{
				Pid:    pid,
				Title:  item,
				Status: true,
			}
			if o.Read(&cate, "title"); cate.Id == 0 {
				_, err = o.Insert(&cate)
			}
		}
	}
	return
}

func (m *Category) UpdateField(id int, field string, value string) (err error) {
	_, err = orm.NewOrm().QueryTable(m.TableName()).Filter("id", id).Update(orm.Params{field: value})
	return
}
