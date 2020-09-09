package controllers

import "fmt"

type CategoryController struct {
	BaseController
}

func (c *CategoryController) Show() {
	fmt.Println("category/show")
	c.TplName = "category/show.html"
}
