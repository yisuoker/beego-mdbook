package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"mdbook/common"
	"mdbook/utils"
	"time"
)

type Member struct {
	MemberId      int       `orm:"pk;auto" json:"member_id"`
	Account       string    `orm:"size(30);unique" json:"account"`
	Nickname      string    `orm:"size(30);unique" json:"nickname"`
	Password      string    ` json:"-"`
	Description   string    `orm:"size(640)" json:"description"`
	Email         string    `orm:"size(100);unique" json:"email"`
	Phone         string    `orm:"size(20);null;default(null)" json:"phone"`
	Avatar        string    `json:"avatar"`
	Role          int       `orm:"default(1)" json:"role"`
	RoleName      string    `orm:"-" json:"role_name"`
	Status        int       `orm:"default(0)" json:"status"`
	CreateTime    time.Time `orm:"type(datetime);auto_now_add" json:"create_time"`
	CreateAt      int       `json:"create_at"`
	LastLoginTime time.Time `orm:"type(datetime);null" json:"last_login_time"`
}

func (m *Member) TableName() string {
	return TNMembers()
}

func NewMember() *Member {
	return &Member{}
}

// crud create, find, update, delete
func (m *Member) Create() error {
	// 验证用户是否存在
	cond := orm.NewCondition().Or("email", m.Email).Or("nickname", m.Nickname).Or("account", m.Account)
	var one Member
	o := orm.NewOrm()
	err := o.QueryTable(m.TableName()).SetCond(cond).One(&one, "member_id", "nickname", "account", "email")
	fmt.Println(err)
	// TODO 查不到rows
	//if err != nil {
	//	return err
	//}
	if one.MemberId > 0 {
		if one.Nickname == m.Nickname {
			return errors.New("昵称已存在")
		}
		if one.Email == m.Email {
			return errors.New("邮箱已存在")
		}
		if one.Account == m.Account {
			return errors.New("用户已存在")
		}
	}

	hash, err := utils.PasswordHash(m.Password)
	if err != nil {
		return err
	}
	m.Password = hash
	_, err = o.Insert(m)
	if err != nil {
		return err
	}

	m.RoleName = common.Role(m.Role)
	return nil
}

func (m *Member) Find(id int) (*Member, error) {
	m.MemberId = id
	if err := orm.NewOrm().Read(m); err != nil {
		return m, err
	}
	m.RoleName = common.Role(m.Role)
	return m, nil
}

func (m *Member) Update(cols ...string) error {
	if _, err := orm.NewOrm().Update(m, cols...); err != nil {
		return err
	}
	return nil
}

func (m *Member) Delete(cols ...string) error {
	// TODO 判断用户权限
	m.Status = 0
	err := m.Update(cols...)
	if err != nil {
		return err
	}
	return nil
}

func (m *Member) Login(account string, password string) (*Member, error) {
	member := &Member{}
	err := orm.NewOrm().QueryTable(m.TableName()).Filter("account", account).Filter("status", 1).One(member)
	if err != nil {
		return member, err
	}

	ok, err := utils.PasswordVerify(member.Password, password)
	if ok && err == nil {
		return member, nil
	} else {
		return member, errors.New("密码错误")
	}
}
