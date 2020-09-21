package routers

import (
	"github.com/astaxie/beego"
	"mdbook/controllers"
)

func init() {
	//beego.Router("/", &controllers.MainController{})

	// test
	beego.Router("/test/:id", &controllers.TestController{}, "*:Detail")

	// home
	beego.Router("/", &controllers.HomeController{}, "get:Index")

	// login, register
	beego.Router("/register", &controllers.AccountController{}, "*:Register")
	beego.Router("/login", &controllers.AccountController{}, "*:Login")
	beego.Router("/logout", &controllers.AccountController{}, "*:Logout")

	// category
	beego.Router("/category/:id", &controllers.CategoryController{}, "get:Show")
	beego.Router("/category/settings", &controllers.CategoryController{}, "get:Settings")
	beego.Router("/category", &controllers.CategoryController{}, "post:Store")
	beego.Router("/category/:id/update", &controllers.CategoryController{}, "post:Update")
	beego.Router("/category/:id/delete", &controllers.CategoryController{}, "get:Delete")
	beego.Router("/category/:id/icon-update", &controllers.CategoryController{}, "get:IconUpdate")

	// books
	// TODO urlfor ?:token 正则路由怎么生成
	beego.Router("/books/:id", &controllers.BooksController{}, "get:Show")

	// members
	beego.Router("/members/:id", &controllers.MembersController{}, "get:Index")
	beego.Router("/members/:id/collection", &controllers.MembersController{}, "get:Collection")
	beego.Router("/members/:id/follow", &controllers.MembersController{}, "get:Follow")
	beego.Router("/members/:id/fans", &controllers.MembersController{}, "get:Fans")
	//beego.Router("/members/:id/follow", &controllers.MembersController{}, "post:FollowUpdate")
	beego.Router("/members/:id/settings", &controllers.MembersController{}, "*:Settings")
	beego.Router("/members/:id/books/?:private", &controllers.MembersController{}, "get:Books")
}
