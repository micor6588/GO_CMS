package models

import "github.com/astaxie/beego/orm"

func init() {
	// orm.RegisterDataBase("default", "mysql", "root@371871@tcp(127.0.0.1:3306)/xcms?charset=utf8")
	orm.RegisterModel(new(MenuModel))
}
