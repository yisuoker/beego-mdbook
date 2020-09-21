package models

import (
	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(
		new(Member),
		new(Category),
		new(Books),
		new(BookCategory),
		new(Documents),
		new(DocumentStore),
		new(Relationship),
	)
}

// 多个数据库怎么获取表前缀
//func TablePrefix() string {
//	return beego.AppConfig.String("")
//}

func TNMembers() string {
	return "md_members"
}

func TNCategory() string {
	return "md_category"
}

func TNBooks() string {
	return "md_books"
}

func TNBookCategory() string {
	return "md_book_category"
}

func TNDocuments() string {
	return "md_documents"
}

func TNDocumentStore() string {
	return "md_document_store"
}

func TNAttachments() string {
	return "md_attachments"
}

func TNRelationship() string {
	return "md_relationship"
}
