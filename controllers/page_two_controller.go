package controllers

import "github.com/astaxie/beego"

type PageTwoController struct {
	beego.Controller
}

func (c *PageTwoController) Get() {
	user := c.GetSession("login")
	if user == nil {
		c.Redirect("/login", 302)
	}

	c.TplName = "pages/pageTwo.html"
}
