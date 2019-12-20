package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about PmsFeightTemplate
type PmsFeightTemplateController struct {
	BaseController
}

// @Title CreatePmsFeightTemplate
// @Description create PmsFeightTemplate
// @Param	body		body 	models.PmsFeightTemplate	true		"body for PmsFeightTemplate content"
// @Success 200 {int} models.PmsFeightTemplate.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PmsFeightTemplateController) Post() {
	var pmsFeightTemplate models.PmsFeightTemplate
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsFeightTemplate)
	uid, _ := models.AddPmsFeightTemplate(pmsFeightTemplate)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all PmsFeightTemplates
// @Success 200 {object} models.PmsFeightTemplate
// @router / [get]
func (u *PmsFeightTemplateController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPmsFeightTemplate(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get PmsFeightTemplate by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PmsFeightTemplate
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PmsFeightTemplateController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsFeightTemplate, err := models.GetPmsFeightTemplate(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, pmsFeightTemplate}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the PmsFeightTemplate
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.PmsFeightTemplate	true		"body for PmsFeightTemplate content"
// @Success 200 {object} models.PmsFeightTemplate
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PmsFeightTemplateController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsFeightTemplate models.PmsFeightTemplate
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsFeightTemplate)
		uu, err := models.UpdatePmsFeightTemplate(uid, &pmsFeightTemplate)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, uu}
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the PmsFeightTemplate
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PmsFeightTemplateController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePmsFeightTemplate(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
