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
	beego.Router("/books/:id/?:token/?:tab", &controllers.BooksController{}, "get:Show")
}
