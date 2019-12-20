package controllers

import (
	"encoding/json"
	"first_go/models"

	"github.com/astaxie/beego"
)

// Operations about Register
type RegisterController struct {
	beego.Controller
}

// @Title Register
// @Description Register umsMember
// @Param	body		body 	models.UmsMember	true		"body for UmsMember content"
// @Success 200 {int} models.UmsMember.Id
// @Failure 403 body is empty
// @router / [post]
func (u *RegisterController) Post() {
	var umsMember models.UmsMember
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &umsMember)
	uid, err := models.RegisterUmsMember(umsMember)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = &ErrResponse{code, message}
	} else {
		u.Data["json"] = &Response{code, message, uid}
	}
	u.ServeJSON()
}
