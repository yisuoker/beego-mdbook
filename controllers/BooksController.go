package controllers

import (
	"fmt"
	"mdbook/common"
	"mdbook/models"
	"strings"
	"time"
)

type BooksController struct {
	BaseController
}

//详情
func (c *BooksController) Show() {
	fmt.Println("books/show")
	fmt.Println("router----------", c.URLFor("BooksController.Show", ":id", 1))

	// id, tab, token
	id, _ := c.GetInt(":id")
	token := c.GetString("token")
	fmt.Println("token----", token)
	if id < 0 {
		c.Abort("参数异常")
	}
	tab := c.GetString("tab")
	if "" == tab {
		tab = "default"
	}

	// 获取图书信息并验证图书权限
	book, err := new(models.Books).Show(id, token)
	if err != nil {
		c.Abort(err.Error())
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

//更新图书
///books/:id => post:/books/update (/book/setting/save => post:/book/save-book)
//
//生成令牌
///books/:id/token => post:/books/TokenUpdate (/book/setting/token => post:/book/create-token)

//收藏
func (c *BooksController) CollectStore() {
	fmt.Println("Books/CollectStore")
	c.TplName = "books/collect_store.html"
}

//保存
func (c *BooksController) Store() {
	fmt.Println("Books/Store")

	BookName := strings.TrimSpace(c.GetString("book_name"))
	Author := strings.TrimSpace(c.GetString("author"))
	AuthorURL := strings.TrimSpace(c.GetString("author_url"))
	Identify := strings.TrimSpace(c.GetString("identify"))
	Description := strings.TrimSpace(c.GetString("description"))
	PrivatelyOwned, _ := c.GetInt("privately_owned")

	if BookName == "" {
		c.JsonResult(1, "请填写图书名称")
	}
	if strings.Count(Description, "") > 500 {
		c.JsonResult(1, "图书描述应小于500字")
	}
	if PrivatelyOwned != 0 || PrivatelyOwned != 1 {
		PrivatelyOwned = 1 // 默认私有
	}

	book := models.Books{
		BookName:       BookName,
		Author:         Author,
		AuthorURL:      AuthorURL,
		Identify:       Identify,
		Description:    Description,
		PrivatelyOwned: PrivatelyOwned,
		CommentCount:   0,
		DocCount:       0,
		Cover:          common.DefaultCover(),
		MemberId:       c.Member.MemberId,
		Editor:         "markdown",
		ReleaseTime:    time.Now(),
		Score:          40,
	}

	if err := book.Create(); err != nil {
		c.JsonResult(1, "操作失败")
	}
	c.JsonResult(0, "操作成功")
}

//编辑
func (c *BooksController) Edit() {
	fmt.Println("Books/Edit")

	// book
	id, _ := c.GetInt(":id")
	book, err := new(models.Books).Get(id)
	if err != nil || book == nil {
		c.Abort("找不到数据")
	}
	c.Data["Model"] = book

	// book category
	if bookCategories, rows, _ := new(models.BookCategory).GetByBookId(id); rows > 0 {
		var maps = make(map[int]bool)
		for _, cate := range bookCategories {
			maps[cate.Id] = true
		}
		c.Data["Maps"] = maps
	}

	// category
	cates, err := new(models.Category).FindAll(-1, -1)
	if err == nil {
		c.Data["Cates"] = cates
	}

	c.TplName = "books/edit.html"
}

//更新
func (c *BooksController) Update() {
	fmt.Println("Books/Update")

	// book
	id, _ := c.GetInt(":id")
	book, err := new(models.Books).Get(id)
	if err != nil || book == nil {
		c.Abort("找不到数据")
	}

	bookName := strings.TrimSpace(c.GetString("book_name"))
	description := strings.TrimSpace(c.GetString("description"))
	if bookName == "" {
		c.JsonResult(1, "请填写图书名称")
	}
	if strings.Count(description, "") > 500 {
		c.JsonResult(1, "图书描述应小于500字")
	}
	book.BookName = bookName
	book.Description = description
	if err := book.Update(); err != nil {
		c.JsonResult(1, "操作失败")
	}

	// category
	if cids, ok := c.Ctx.Request.Form["cid"]; len(cids) > 0 && ok {
		fmt.Println("--------- cids: ", cids)
		new(models.BookCategory).UpdateBookCategories(id, cids)
	}

	c.JsonResult(0, "ok", book)
}

//更新令牌
func (c *BooksController) TokenUpdate() {
	fmt.Println("Books/TokenUpdate")
	c.TplName = "books/token_update.html"
}

//发布
func (c *BooksController) Release() {
	fmt.Println("Books/Release")
	c.TplName = "books/release.html"
}
