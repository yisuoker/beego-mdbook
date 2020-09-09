package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mdbook/models"
)

type HomeController struct {
	BaseController
}

func (c *HomeController) Index() {
	// 获取所有分类
	if cates, err := new(models.Category).FindAll(-1, 1); err == nil {
		fmt.Println(cates)
		c.Data["Cates"] = cates
	} else {
		beego.Error(err.Error())
	}
	c.TplName = "home/index.html"
}
