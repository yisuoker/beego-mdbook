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

	// books
	beego.Router("/books/:id", &controllers.BooksController{}, "get:Show")
}
