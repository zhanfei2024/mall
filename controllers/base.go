package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
)

// BaseController 结构体
type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	phone := c.GetSession("phone")
	password := c.GetSession("password")
	if phone != nil && password != nil && phone == c.Ctx.GetCookie("phone") && password == c.Ctx.GetCookie("password") {
		c.Data["phone"] = phone
		c.Data["password"] = password
		fmt.Println(phone, password, c.Ctx.GetCookie("phone"), c.Ctx.GetCookie("password"))
	} else {
		fmt.Println(phone, password, c.Ctx.GetCookie("phone"), c.Ctx.GetCookie("password"))
		code, message := DecodeErr(ErrNoPermission)
		c.Data["json"] = ErrResponse{code, message}
		c.ServeJSON()
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
