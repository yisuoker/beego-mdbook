package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mdbook/common"
	"mdbook/models"
	"mdbook/utils"
	"strings"
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

func (c *CategoryController) Settings() {
	fmt.Println("category/settings")
	fmt.Println(c.URLFor("CategoryController.Update", ":id", 1))

	// 获取所有的分类
	cates, err := new(models.Category).FindAll(-1, -1)
	if err != nil {
		beego.Error(err.Error())
	}

	// 初始化icon
	var parents []models.Category
	for idx, item := range cates {
		if "" == strings.TrimSpace(item.Icon) {
			item.Icon = common.DefaultBookIcon()
		} else {
			item.Icon = utils.ShowImg(item.Icon)
		}
		if 0 == item.Pid {
			parents = append(parents, item)
		}
		cates[idx] = item
	}

	c.Data["Parents"] = parents
	c.Data["Cates"] = cates
	c.Data["IsCategory"] = true
	c.TplName = "category/settings.html"
}

func (c *CategoryController) Store() {
	fmt.Println("category/Store")
	//pid=3&cates=vue2

	// ajax json
	pid, _ := c.GetInt("pid")                        // 默认0
	cates := strings.TrimSpace(c.GetString("cates")) // 默认""
	if 0 == len(cates) {
		c.JsonResult(2, "参数异常")
	}
	fmt.Println("pid=", pid, "cates=", cates)
	if err := new(models.Category).MultiCreate(pid, cates); err != nil {
		c.JsonResult(1, "新增失败"+err.Error())
	}
	c.JsonResult(0, "新增成功")
}

func (c *CategoryController) Update() {
	fmt.Println("category/Update")
	//manager/update-cate?id=15&field=intro&value=vue

	// 每次只更新一个字段信息
	id, _ := c.GetInt(":id")
	if id <= 0 {
		c.JsonResult(2, "参数异常")
	}
	field := strings.TrimSpace(c.GetString("field"))
	value := strings.TrimSpace(c.GetString("value"))
	fmt.Println("field=", field, "value=", value)
	if err := new(models.Category).UpdateField(id, field, value); err != nil {
		c.JsonResult(1, "更新失败"+err.Error())
	}
	c.JsonResult(0, "更新成功")
}

func (c *CategoryController) Delete() {
	fmt.Println("category/Delete")
	// 逻辑删除 update(status=2)
	c.TplName = "category/Delete.html"
}

func (c *CategoryController) IconUpdate() {
	fmt.Println("category/IconUpdate")
	c.TplName = "category/IconUpdate.html"
}
