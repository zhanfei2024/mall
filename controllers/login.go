package controllers

import (
	"encoding/json"
	"mall/models"

	"github.com/astaxie/beego"
)

// Operations about Login
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description Login umsMember
// @Param	body		body 	models.UmsMember	true		"body for UmsMember content"
// @Success 200 {object} models.UmsMember
// @Failure 403 body is empty
// @router / [post]
func (u *LoginController) Post() {
	var umsMember models.UmsMember
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &umsMember)

	umsMemberRes, err := models.LoginOfUmsMember(umsMember)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.SetSession("phone", umsMember.Phone)
		u.SetSession("password", umsMember.Password)
		u.TplName = "index.tpl"

		u.Ctx.SetCookie("phone", umsMember.Phone, 36000, "/")
		u.Ctx.SetCookie("password", umsMember.Password, 36000, "/")

		u.Data["json"] = Response{code, message, umsMemberRes}
	}
	u.ServeJSON()

}
