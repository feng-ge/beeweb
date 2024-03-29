package controllers

import (
	"beeweb/models"

	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Get() {
	// c.Data["Copyby"] = "copy right fengge3"
	this.Data["IsIndex"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "index.html"
	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}
