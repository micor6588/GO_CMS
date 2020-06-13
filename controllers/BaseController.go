package controllers

import (
	"GO_CMS/models"
	"fmt"
	"strings"

	"github.com/astaxie/beego"
)

// BaseController 基本控制器
type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
}

func (c *BaseController) Prepare() {
	//赋值
	c.controllerName, c.actionName = c.GetControllerAndAction()
	beego.Informational(c.controllerName, c.actionName)

	// TODO: 保存用户数据
	fmt.Println("beego: prepare: " + c.controllerName + "," + c.actionName)
	c.Data["Menu"] = models.MenuStruct()
}

// 设置模板
// 第一个参数模板，第二个参数为layout
func (c *BaseController) setTpl(template ...string) {
	var tplName string
	layout := "/views/common/layout.html"
	switch {
	case len(template) == 1:
		tplName = template[0]
	case len(template) == 2:
		tplName = template[1]
	default:
		// 不要"Controller” 这10个字母
		ctrlName := strings.ToLower(c.controllerName[0 : len(c.controllerName)-10])
		actionName := strings.ToLower(c.actionName)
		tplName = ctrlName + "/" + actionName + ".html"
	}
	_, found := c.Data["Footer"]
	if !found {
		c.Data["Footer"] = "/views/menu/footerjs.html"
	}
	c.Layout = layout
	c.TplName = tplName
}
