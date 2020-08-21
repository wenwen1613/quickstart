package controllers

import (
	"github.com/astaxie/beego"
)

type Result struct {
	Code int
	Msg  string
	Data interface{}
}
type User struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
}

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.TplName = "index.tpl"
}

func (c *UserController) Info() {
	c.Ctx.WriteString("User-Info")
}

func (c *UserController) List() {
	c.Ctx.WriteString("user list")
}

func (c *UserController) Add() {
	u := User{}
	userErr := c.ParseForm(&u)
	if userErr != nil {
		data := &Result{
			Code: 1001,
			Msg:  "参数错误",
			Data: make(map[string]string),
		}
		c.Data["json"] = data
		c.ServeJSON()
	} else {
		data := &Result{
			Code: 1000,
			Msg:  "OK",
			Data: &u,
		}
		c.Data["json"] = data
		c.ServeJSON()
	}
}
