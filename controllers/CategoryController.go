package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mdbook/models"
)

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Show() {
	fmt.Println("category/show")

	//category
	id, _ := c.GetInt(":id")
	if id > 0 {
		cate, _ := new(models.Category).Find(id)
		fmt.Println("cate--------", cate)
		c.Data["Cate"] = cate
	} else {
		beego.Error("参数异常")
		c.Abort("404")
	}

	//books
	books, err := new(models.Books).GetByCategoryId(id)
	if err != nil {
		beego.Error(err)
		c.Abort("404")
	}
	c.Data["Lists"] = books
	c.Data["Cid"] = id

	c.TplName = "category/show.html"
}
