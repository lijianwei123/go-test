package controllers

import (
	"github.com/astaxie/beego"
	"github.com/lijianwei123/go-test/shortlink/work"
	"strconv"
)

type GitController struct {
	beego.Controller
}

func (this *GitController) Pull() {
	//	c.Data["Website"] = "beego.me"
	//	c.Data["Email"] = "astaxie@gmail.com"
	//	c.TplNames = "index.tpl"
	success := work.Pull()
	this.Ctx.WriteString(strconv.FormatBool(success))
}
