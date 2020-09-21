package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mdbook/models"
)

type BooksController struct {
	BaseController
}

func (c *BooksController) Show() {
	fmt.Println("books/show")
	fmt.Println("router----------", c.URLFor("BooksController.Show", ":id", 1))

	// id, tab, token
	id, _ := c.GetInt(":id")
	token := c.GetString("token")
	fmt.Println("token----", token)
	if id < 0 {
		beego.Error("参数异常")
	}
	tab := c.GetString("tab")
	if "" == tab {
		tab = "default"
	}

	// 获取图书信息并验证图书权限
	book, err := new(models.Books).Show(id, token)
	if err != nil {
		beego.Error(err.Error())
	}
	c.Data["Book"] = book

	// 目录
	c.Data["Menu"], _ = new(models.Documents).GetByBookId(id)

	// 评论
	//c.Data["Comments"]
	//c.Data["MyScore"]

	c.Data["Tab"] = tab
	c.Data["Token"] = token
	c.TplName = "books/show.html"
}
