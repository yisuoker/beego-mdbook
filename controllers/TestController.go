package controllers

import (
	"github.com/astaxie/beego"
	"mdbook/models"
)

type TestController struct {
	beego.Controller
}

func (c *TestController) Index() {
	// model

	// tpl
}

func (c *TestController) Detail()  {
	id, _:= c.GetInt(":id")

	test := new(models.Test)
	res := test.GetById(id)

	c.Data["res"] = res
	c.TplName = "test/detail.html"
}
