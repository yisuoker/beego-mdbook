package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Documents struct {
	DocumentId   int            `orm:"pk;auto;column(document_id)" json:"doc_id"`
	DocumentName string         `orm:"column(document_name);size(500)" json:"doc_name"`
	Identify     string         `orm:"column(identify);size(100);index;null;default(null)" json:"identify"`
	BookId       int            `orm:"column(book_id);type(int)" json:"book_id"`
	ParentId     int            `orm:"column(parent_id);type(int);default(0)" json:"parent_id"`
	OrderSort    int            `orm:"column(order_sort);default(0);type(int)" json:"order_sort"`
	Release      string         `orm:"column(release);type(text);null" json:"release"`
	CreateTime   time.Time      `orm:"column(create_time);type(datetime);auto_now_add" json:"create_time"`
	MemberId     int            `orm:"column(member_id);type(int)" json:"member_id"`
	ModifyTime   time.Time      `orm:"column(modify_time);type(datetime);default(null);auto_now" json:"modify_time"`
	ModifyAt     int            `orm:"column(modify_at);type(int)" json:"-"`
	Version      int64          `orm:"type(bigint);column(version)" json:"version"`
	AttachList   []*Attachments `orm:"-" json:"attach"`
	Vcnt         int            `orm:"column(vcnt);default(0)" json:"vcnt"`
	Markdown     string         `orm:"-" json:"markdown"`
}

func (m *Documents) TableName() string {
	return TNDocuments()
}

func (m *Documents) GetByBookId(bookId int) (docs []*Documents, err error) {
	var docsAll []*Documents
	o := orm.NewOrm()
	cols := []string{"document_id", "document_name", "member_id", "parent_id", "book_id", "identify"}
	fmt.Println("---------------start")
	_, err = o.QueryTable(m.TableName()).Filter("book_id", bookId).Filter("parent_id", 0).OrderBy("order_sort", "document_id").Limit(5000).All(&docsAll, cols...)
	fmt.Println("---------------end")
	for _, doc := range docsAll {
		docs = append(docs, doc)
	}
	return
}
