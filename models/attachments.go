package models

import "time"

type Attachments struct {
	AttachmentId int `orm:"pk;auto" json:"attachment_id"`
	BookId       int ` json:"book_id"`
	DocumentId   int ` json:"doc_id"`
	Name         string
	Path         string    `orm:"size(2000)" json:"file_path"`
	Size         float64   `orm:"type(float)" json:"file_size"`
	Ext          string    `orm:"size(50)" json:"file_ext"`
	HttpPath     string    `orm:"size(2000)" json:"http_path"`
	CreateTime   time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	CreateAt     int       `orm:"type(int)" json:"create_at"`
}

//orm回调TableName
func (m *Attachments) TableName() string {
	return TNAttachments()
}
