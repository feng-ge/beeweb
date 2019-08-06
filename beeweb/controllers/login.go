package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (this *LoginController) Post() {
	username := this.Input().Get("username")
	password := this.Input().Get("password")
	autoLogin := this.Input().Get("autoLogin") == "on"

	if beego.AppConfig.String("username")==username && beego.AppConfig.String("password")==password{
		maxAge := 0
		if autoLogin{
			maxAge=1<<31-1
		}
		this.Ctx.SetCookie("username",username,maxAge,"/")
		this.Ctx.SetCookie("password",password,maxAge,"/")
	}
	this.Redirect("/",301)
	return
}