package main

import (
	"beeweb/controllers"
	"beeweb/models"
	_ "beeweb/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func init() {
	//注册数据库
	models.RegisterDB()
}

func main() {
	//开启ORM调试模式
	orm.Debug = true
	orm.RunSyncdb("default", false, true)
	//
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.Run()
}
