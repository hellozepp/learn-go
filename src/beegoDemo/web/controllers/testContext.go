package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
)

type TestContextController struct {
	beego.Controller
}

func (c *TestContextController) Get() {
	c.Ctx.WriteString(c.Ctx.Input.IP() + ":" + strconv.Itoa(c.Ctx.Input.Port()))

	c.Ctx.WriteString(c.Ctx.Input.Query("name")) //等价于php中的 $_REQUEST["name"]

	m := make(map[string]float64)
	m["zhangsan"] = 98.7
	c.Ctx.Output.JSON(m, false, false)
}
