package models

import "github.com/astaxie/beego/orm"

type MenuModel struct {
	Mid    int `orm:"pk;auto"`
	Parent int
	Name   string `orm:"size(45)"`
	Seq    int
	Format string `orm:"size(2048);default({})"`
}
type MenuTree struct {
	MenuModel
	Child []MenuModel
}

func (m *MenuModel) TableName() string {
	return "xcms_menu"
}

func MenuStruct() map[int]MenuTree {
	query := orm.NewOrm().QueryTable("xcms_menu")
	data := make([]*MenuModel, 0)
	query.OrderBy("parnt", "-seq").All(&data)
	var menu = make(map[int]MenuTree)
	if len(data) > 0 {
		for _, v := range data {
			if v.Parent == 0 {
				var tree = new(MenuTree)
				tree.MenuModel = *v
				menu[v.Mid] = *tree
			} else {
				if temp, ok := menu[v.Parent]; ok {
					temp.Child = append(temp.Child, *v)
					menu[v.Parent] = temp
				}
			}
		}
	}
	return menu
}

// func MenuList() ([]*MenuModel, int64) {

// }
