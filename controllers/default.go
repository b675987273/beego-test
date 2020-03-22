package controllers

import (
	"fmt"
	model "go-pratice-demo/beego-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "test.html"
	// o := orm.NewOrm()

	// user := model.User{1, "AA", "bb"}

	// i, err := o.Insert(&user)

	// if err != nil {
	// 	fmt.Println(err, "  ", i)
	// 	return
	// }
	o := orm.NewOrm()
	var u model.User
	u.Id = 1
	err := o.Read(&u, "Id")
	if err != nil {
		fmt.Println(err)
		return
	}
	u.Id = 1
	u.Name = "cc"
	i, err := o.Update(&u)
	if err != nil {
		fmt.Println(err, "  ", i)
		return
	}
	fmt.Println(u)
	i, err = o.Delete(&u)
	if err != nil {
		fmt.Println(err, "  ", i)
		return
	}

}
func (c *MainController) Post() {
	uname := c.GetString("un")
	pwd := c.GetString("pwd")

	if uname == "" || pwd == "" {
		c.Redirect("/reg", 404)
		beego.Info("kong\n")
		return
	}
}
