package models

type Relationship struct {
	RelationshipId int `orm:"pk;auto;" json:"relationship_id"`
	MemberId       int `json:"member_id"`
	BookId         int ` json:"book_id"`
	RoleId         int `json:"role_id"` // common.BookRole
}

func (m *Relationship) TableName() string {
	return TNRelationship()
}

// orm回调 联合唯一索引
func (m *Relationship) TableUnique() [][]string {
	return [][]string{
		[]string{"MemberId", "BookId"},
	}
}

func NewRelationship() *Relationship {
	return &Relationship{}
}
