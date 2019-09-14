package controllers

import "github.com/astaxie/beego"

type PageOneController struct {
	beego.Controller
}

func (c *PageOneController) Get() {
	user := c.GetSession("login")
	if user == nil {
		c.Redirect("/login", 302)
	}

	c.TplName = "pages/pageOne.html"
}
