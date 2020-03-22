package main

import (
	"go-pratice-demo/beego-test/controllers"
	_ "go-pratice-demo/beego-test/models"
	_ "go-pratice-demo/beego-test/routers"

	"github.com/astaxie/beego"
)

func main() {
	controllers.InsertUser(3)
	beego.SetStaticPath("d1", "static/img")
	beego.Run()
}
