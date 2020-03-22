package model

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Pwd  string
	Ord  *Order `orm:"rel(fk)"`
}
type Order struct {
	Id   int
	Dta  string
	User []*User `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:woshi5200@tcp(127.0.0.1:3306)/beego?charset=utf8")
	orm.RegisterModel(new(User), new(Order))
	orm.RunSyncdb("default", false, true)
}
