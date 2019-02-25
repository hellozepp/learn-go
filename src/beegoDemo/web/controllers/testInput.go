package controllers

import (
	"github.com/astaxie/beego"
)

type TestInputController struct {
	beego.Controller
}

type User struct {
	Username string
	Password string
}

func (c *TestInputController) Get() {
	//id := c.GetString("id")
	//c.Ctx.WriteString("<html>" + id + "<br/>")

	//name := c.Input().Get("name")
	//c.Ctx.WriteString(name + "</html>")
	name := c.GetSession("name") //return interface{}
	pass := c.GetSession("password")
	if nameString, ok := name.(string); ok && nameString != "" {
		c.Ctx.WriteString("name:" + nameString + "password:" + pass.(string))
	} else {
		c.Ctx.WriteString(`<html><form action="http://127.0.0.1:8080/test_input" method="post"> 
							<input type="text" name="Username"/>
							<input type="password" name="Password"/>
							<input type="submit" value="提交"/>
					   </form></html>`)
	}
}

func (c *TestInputController) Post() {
	u := User{}
	if err := c.ParseForm(&u); err != nil {
		//process error
	}

	c.Ctx.WriteString("Username:" + u.Username + " Password:" + u.Password)
}
