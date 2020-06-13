package controllers

// MenuController 基本控制器
type MenuController struct {
	BaseController
}

func (c *MenuController) Index() {
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["footjs"] = "/views/menu/footerjs.html"
	c.setTpl("/views/menu/index.html")
}
func (c *MenuController) List() {

}

func (c *MenuController) Add() {

}

func (c *MenuController) AddDo() {

}
func (c *MenuController) Edit() {

}
func (c *MenuController) EditDo() {

}
