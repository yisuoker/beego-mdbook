package models

type DocumentStore struct {
	DocumentId int    `orm:"pk;auto;column(document_id)"`
	Markdown   string `orm:"type(text);"` //markdown内容
	Content    string `orm:"type(text);"` //html内容
}

func (m *DocumentStore) TableName() string {
	return TNDocumentStore()
}
