package routers

import (
	"mdbook/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    // test
    beego.Router("/test/:id", &controllers.TestController{}, "*:Detail")
}
