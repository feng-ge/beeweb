package controllers

import (
	"beeweb/models"

	"github.com/astaxie/beego"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplName = "topic.html"
	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	tid := this.Input().Get("tid")
	category := this.Input().Get("category")
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//查分类
	var err error
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.TplName = "topic_view.html"
	topic, err := models.GetTopic(this.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = this.Ctx.Input.Param("0")
	this.Data["IsLogin"] = checkAccount(this.Ctx)
}

func (this *TopicController) Modify() {
	this.TplName = "topic_modify.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}
	this.Data["Topic"] = topic
	this.Data["Tid"] = tid
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	//查分类
	this.Data["Categories"], err = models.GetAllCategories()
	if err != nil {
		beego.Error(err)
	}
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}
	err := models.DeleteTopic(this.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	this.Redirect("/topic", 302)
}
