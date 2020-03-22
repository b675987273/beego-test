package controllers

import (
	"fmt"
	model "go-pratice-demo/beego-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type UserController struct {
	beego.Controller
}

func InsertUser(id int) {
	o := orm.NewOrm()
	u := model.User{id, "aa", "dd", nil}
	x, err := o.Insert(&u)
	if err != nil {
		fmt.Println(err, "1", x)
		return
	}
	beego.Info("success1:", x)
}
func QueryUser(id int) {
	o := orm.NewOrm()
	u := model.User{Id: id, Name: "ss"}
	x, err := o.Update(&u)
	if err != nil {
		fmt.Println(err, "2", x)
		return
	}
	beego.Info("success2:", x)
}
func InsertOrder(id int) {
	o := orm.NewOrm()
	or := model.Order{id, "??", nil}
	_, err := o.Insert(&or)
	if err != nil {
		fmt.Println(err, "3")
		return
	}

}

func (u *UserController) Get() {
	u.Ctx.WriteString("test\n")
}
func (u *UserController) GetInfo() {
	id := u.Ctx.Input.Param(":id")
	u.Ctx.WriteString(id)

}
func (u *UserController) Post() {
	u.Ctx.WriteString("Post")
}
