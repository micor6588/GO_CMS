package main

import (
	_ "GO_CMS/routers"
	_ "GO_CMS/sysinit"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
