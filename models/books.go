package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mdbook/utils"
	"strconv"
	"strings"
	"time"
)

// 图书
type Books struct {
	BookId         int       `orm:"pk;auto" json:"book_id"`
	BookName       string    `orm:"size(500)" json:"book_name"`       //名称
	Identify       string    `orm:"size(100);unique" json:"identify"` //唯一标识
	OrderIndex     int       `orm:"default(0)" json:"order_index"`
	Description    string    `orm:"size(1000)" json:"description"`       //图书描述
	Cover          string    `orm:"size(1000)" json:"cover"`             //封面地址
	Editor         string    `orm:"size(50)" json:"editor"`              //编辑器类型: "markdown"
	Status         int       `orm:"default(0)" json:"status"`            //状态:0 正常 ; 1 已删除
	PrivatelyOwned int       `orm:"default(0)" json:"privately_owned"`   // 是否私有: 0 公开 ; 1 私有
	PrivateToken   string    `orm:"size(500);null" json:"private_token"` // 私有图书访问Token
	MemberId       int       `orm:"size(100)" json:"member_id"`
	CreateTime     time.Time `orm:"type(datetime);auto_now_add" json:"create_time"` //创建时间
	ModifyTime     time.Time `orm:"type(datetime);auto_now_add" json:"modify_time"`
	ReleaseTime    time.Time `orm:"type(datetime);" json:"release_time"` //发布时间
	DocCount       int       `json:"doc_count"`                          //文档数量
	CommentCount   int       `orm:"type(int)" json:"comment_count"`
	Vcnt           int       `orm:"default(0)" json:"vcnt"`              //阅读次数
	Collection     int       `orm:"column(star);default(0)" json:"star"` //收藏次数
	Score          int       `orm:"default(40)" json:"score"`            //评分
	CntScore       int       //评分人数
	CntComment     int       //评论人数
	Author         string    `orm:"size(50)"`                      //来源
	AuthorURL      string    `orm:"column(author_url);size(1000)"` //来源链接
}

// 获取表名
func (m *Books) TableName() string {
	return TNBooks()
}

// 根据分类id获取图书列表
func (m *Books) GetByCategoryId(categoryId int, fields ...string) (books []Books, err error) {
	if 0 == len(fields) {
		fields = append(fields, "book_id", "book_name", "identify", "cover", "order_index", "private_token")
	}
	fieldStr := "b." + strings.Join(fields, ",b.")
	sqlFmt := "select %v from " + m.TableName() + " b left join " + new(BookCategory).TableName() + " c on c.book_id=b.book_id where c.category_id=" + strconv.Itoa(categoryId)
	sql := fmt.Sprintf(sqlFmt, fieldStr)
	fmt.Println(sql)
	if _, err := orm.NewOrm().Raw(sql).QueryRows(&books); err != nil {
		return books, err
	}
	return
}

// 根据field获取图书
func (m *Books) GetByField(field string, value interface{}, cols ...string) (book *Books, err error) {
	// TODO 错误
	//if 0 == len(cols) {
	//	err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(book)
	//} else {
	//	err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(book, cols...)
	//}
	//return book, err
	if 0 == len(cols) {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m)
	} else {
		err = orm.NewOrm().QueryTable(m.TableName()).Filter(field, value).One(m, cols...)
	}
	return m, err
}

// 图书首页
func (m *Books) Show(id int, token string) (bd *BookData, err error) {
	// book
	//var book *Books
	book := &Books{}
	book, err = m.GetByField("book_id", id)
	fmt.Println("book -----", book)
	if err != nil {
		return bd, err
	}

	// verify：1，超管；2，token
	// TODO 超级管理员是0
	fmt.Println("PrivatelyOwned ---------")
	if 1 == book.PrivatelyOwned {
		if "" != token && strings.EqualFold(token, book.PrivateToken) {

		} else {
			return bd, errors.New("没有权限访问")
		}
	}

	bd = &BookData{
		BookId:         book.BookId,
		BookName:       book.BookName,
		Identify:       book.Identify,
		OrderIndex:     book.OrderIndex,
		Description:    strings.Replace(book.Description, "\r\n", "<br/>", -1),
		PrivatelyOwned: book.PrivatelyOwned,
		PrivateToken:   book.PrivateToken,
		DocCount:       book.DocCount,
		CommentCount:   book.CommentCount,
		CreateTime:     book.CreateTime,
		ModifyTime:     book.ModifyTime,
		Cover:          book.Cover,
		MemberId:       book.MemberId,
		Status:         book.Status,
		Editor:         book.Editor,
		Vcnt:           book.Vcnt,
		Collection:     book.Collection,
		Score:          book.Score,
		ScoreFloat:     utils.ScoreFloat(book.Score),
		CntScore:       book.CntScore,
		CntComment:     book.CntComment,
		Author:         book.Author,
		AuthorURL:      book.AuthorURL,
	}
	return
}

// 图书首页图书信息
type BookData struct {
	BookId         int       `json:"book_id"`
	BookName       string    `json:"book_name"`
	Identify       string    `json:"identify"`
	OrderIndex     int       `json:"order_index"`
	Description    string    `json:"description"`
	PrivatelyOwned int       `json:"privately_owned"`
	PrivateToken   string    `json:"private_token"`
	DocCount       int       `json:"doc_count"`
	CommentCount   int       `json:"comment_count"`
	CreateTime     time.Time `json:"create_time"`
	CreateName     string    `json:"create_name"`
	ModifyTime     time.Time `json:"modify_time"`
	Cover          string    `json:"cover"`
	MemberId       int       `json:"member_id"`
	Username       int       `json:"user_name"`
	Editor         string    `json:"editor"`
	RelationshipId int       `json:"relationship_id"`
	RoleId         int       `json:"role_id"`
	RoleName       string    `json:"role_name"`
	Status         int
	Vcnt           int    `json:"vcnt"`
	Collection     int    `json:"star"`
	Score          int    `json:"score"`
	CntComment     int    `json:"cnt_comment"`
	CntScore       int    `json:"cnt_score"`
	ScoreFloat     string `json:"score_float"`
	LastModifyText string `json:"last_modify_text"`
	Author         string `json:"author"`
	AuthorURL      string `json:"author_url"`
}
