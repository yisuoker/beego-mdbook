package controllers

import "fmt"

type BooksController struct {
	BaseController
}

func (c *BooksController) Show() {
	fmt.Println("books/show")
	c.TplName = "books/show.html"
}
