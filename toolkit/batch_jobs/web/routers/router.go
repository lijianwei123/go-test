package routers

import (
	"github.com/astaxie/beego"
	"github.com/lijianwei123/go-test/toolkit/batch_jobs/web/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
