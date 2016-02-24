package routers

import (
	"github.com/astaxie/beego"
	"github.com/lijianwei123/go-test/shortlink/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.GitController{})
}
