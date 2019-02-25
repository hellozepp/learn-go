package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type TestHttpLibController struct {
	beego.Controller
}

func (c *TestHttpLibController) Get() {
	req := httplib.Get("http://www.baidu.com")
	str, err := req.String()

	if err != nil {
		panic(err)
	}

	c.Ctx.WriteString(str)
}
