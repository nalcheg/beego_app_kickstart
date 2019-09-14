package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"github.com/nalcheg/beego_app_kickstart/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	user := c.GetSession("login")
	if user != nil {
		c.Redirect("/", 302)
	}

	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	var login, password string
	err := c.Ctx.Input.Bind(&login, "login")
	if err != nil {
		log.Printf("%v", err)
		c.Redirect("/", 302)
	}
	err = c.Ctx.Input.Bind(&password, "password")
	if err != nil {
		log.Printf("%v", err)
		c.Redirect("/", 302)
	}
	result := models.CheckLoginPassword(login, password)

	if result == true {
		c.SetSession("login", login)
		c.Redirect("/", 302)
	}

	c.Redirect("/login", 302)
}
