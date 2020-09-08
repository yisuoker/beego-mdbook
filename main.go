package main

import (
	"github.com/astaxie/beego"
	_ "mdbook/routers"
	_ "mdbook/sysinit" // 调用 sysinit/init() 方法
)

func main() {
	beego.Run()
}

