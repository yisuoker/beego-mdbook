package models

import (
	"github.com/astaxie/beego/orm"
	"strconv"
)

type BookCategory struct {
	Id         int //自增主键
	BookId     int //书籍id
	CategoryId int //分类id
}

func (m *BookCategory) TableName() string {
	return TNBookCategory()
}

func (m *BookCategory) GetByBookId(bookId int) (cates []*Category, rows int64, err error) {
	sql := "select c.* " +
		"from md_book_category bc " +
		"left join md_category c on c.id=bc.category_id " +
		"where bc.book_id=?"
	rows, err = orm.NewOrm().Raw(sql, bookId).QueryRows(&cates)
	if err != nil {
		return nil, 0, err
	}
	return
}

func (m *BookCategory) UpdateBookCategories(id int, cids []string) {
	// 先删除再更新
	orm.NewOrm().QueryTable(m.TableName()).Filter("book_id", id).Delete()

	// multi insert
	var bookCates []BookCategory
	for _, cid := range cids {
		cidNum, _ := strconv.Atoi(cid)
		bookCate := BookCategory{
			BookId:     id,
			CategoryId: cidNum,
		}
		bookCates = append(bookCates, bookCate)
	}
	if l := len(bookCates); l > 0 {
		orm.NewOrm().InsertMulti(l, &bookCates)
	}
	// TODO 更新分类统计
}
