package routers

import (
	"github.com/lijianwei123/go-test/shortlink/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
