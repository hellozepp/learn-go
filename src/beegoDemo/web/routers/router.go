package routers

import (
	"beegoDemo/web/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/test", &controllers.TestController{}, "get:Get;post:Post")
	beego.Router("/test_input", &controllers.TestInputController{}, "get:Get;post:Post")
	beego.Router("/test_login", &controllers.TestLoginController{}, "get:Login;post:Post")
	beego.Router("/test_view", &controllers.TestViewController{}, "get:Get;post:Post")
	beego.Router("/test_httplib", &controllers.TestHttpLibController{}, "get:Get;post:Post")
	beego.Router("/test_context", &controllers.TestContextController{}, "get:Get;post:Post")
}
