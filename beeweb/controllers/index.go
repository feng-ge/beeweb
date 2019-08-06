package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (c *IndexController) Get() {
	// c.Data["Copyby"] = "copy right fengge3"
	c.Data["IsIndex"] = true
	c.TplName = "index.html"
}
