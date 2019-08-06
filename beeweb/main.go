package main

import (
	"beeweb/controllers"
	"beeweb/models"
	"github.com/astaxie/beego/orm"
	_ "beeweb/routers"
	"github.com/astaxie/beego"
)

func init()  {
	//注册数据库
	models.RegisterDB()
}

func main() {
	//开启ORM调试模式	
	orm.Debug=true
	orm.RunSyncdb("default",false,true)
	//
	beego.Router("/",&controllers.IndexController{})
	beego.Router("/login",&controllers.LoginController{})
	beego.Run()
}

