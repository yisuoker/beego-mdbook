package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"mdbook/models"
)

type MembersController struct {
	BaseController
	User *models.Member
}

func (c *MembersController) Prepare() {
	c.BaseController.Prepare()

	// 用户信息
	id, _ := c.GetInt(":id")
	if id < 0 {
		c.Abort("参数异常")
	}
	user, err := new(models.Member).Find(id)
	if err != nil {
		c.Abort(err.Error())
	}
	if user == nil {
		c.Abort("用户不存在")
	}

	c.User = user
	c.Data["User"] = user
	c.Data["IsSelf"] = user.MemberId == c.Member.MemberId
	c.Data["Tab"] = "share"
}

func (c *MembersController) Index() {
	fmt.Println("members/index")

	// books
	books, err := new(models.Books).GetByUserId(c.User.MemberId, -1)
	if err != nil {
		c.Abort(err.Error())
	}
	c.Data["Books"] = books
	c.TplName = "members/index.html"
}

func (c *MembersController) Collection() {
	fmt.Println("members/Collection")
	c.TplName = "members/collection.html"
}

func (c *MembersController) Follow() {
	fmt.Println("members/Follow")
	c.TplName = "members/follow.html"
}

func (c *MembersController) Fans() {
	fmt.Println("members/Fans")
	c.TplName = "members/fans.html"
}

func (c *MembersController) FollowUpdate() {
	fmt.Println("members/FollowUpdate")
	c.TplName = "members/follow_update.html"
}

func (c *MembersController) Settings() {
	fmt.Println("members/Settings")
	c.TplName = "members/settings.html"
}

func (c *MembersController) Books() {
	fmt.Println("members/Books")
	private, _ := c.GetInt(":private", 1) // 默认私有图书

	// books
	books, err := new(models.Books).GetByUserId(c.User.MemberId, private)
	if err != nil {
		c.Abort(err.Error())
	}
	b, err := json.Marshal(books)
	if err != nil || len(books) <= 0 {
		c.Data["Result"] = template.JS("[]")
	} else {
		c.Data["Result"] = template.JS(string(b))
	}
	c.Data["Private"] = private
	c.TplName = "members/books.html"
}
