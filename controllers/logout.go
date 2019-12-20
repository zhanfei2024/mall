package controllers

import (
	"encoding/json"
	"mall/models"
)

// Operations about Logout
type LogoutController struct {
	BaseController
}

// @Title Logout
// @Description Logout umsMember
// @Param	body		body 	models.UmsMember	true		"body for UmsMember content"
// @Success 200 {object} models.UmsMember
// @Failure 403 body is empty
// @router / [post]
func (u *LogoutController) Post() {
	var umsMember models.UmsMember
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &umsMember)

	umsMemberRes, err := models.LogoutOfUmsMember(umsMember)

	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.DelSession("phone")
		u.DelSession("password")
		u.Ctx.SetCookie("phone", "", -1)
		u.Ctx.SetCookie("password", "", -1)
		u.Data["json"] = Response{code, message, umsMemberRes}
	}
	u.ServeJSON()

}
