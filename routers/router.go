package routers

import (
	"go-pratice-demo/beego-test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.UserController{})
	beego.Router("/reg", &controllers.MainController{})
}
