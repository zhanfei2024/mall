package controllers

import (
	"encoding/json"
	"first_go/models"
	"first_go/utils"
	"strconv"
)

// Operations about UmsMemberLevel
type UmsMemberLevelController struct {
	BaseController
}

// @Title CreateUmsMemberLevel
// @Description create UmsMemberLevel
// @Param	body		body 	models.UmsMemberLevel	true		"body for UmsMemberLevel content"
// @Success 200 {int} models.UmsMemberLevel.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UmsMemberLevelController) Post() {
	var pmsBrand models.UmsMemberLevel
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
	uid, _ := models.AddUmsMemberLevel(pmsBrand)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all UmsMemberLevels
// @Success 200 {object} models.UmsMemberLevel
// @router / [get]
func (u *UmsMemberLevelController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllUmsMemberLevel(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get UmsMemberLevel by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UmsMemberLevel
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UmsMemberLevelController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsBrand, err := models.GetUmsMemberLevel(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, pmsBrand}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the UmsMemberLevel
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.UmsMemberLevel	true		"body for UmsMemberLevel content"
// @Success 200 {object} models.UmsMemberLevel
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UmsMemberLevelController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsBrand models.UmsMemberLevel
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
		uu, err := models.UpdateUmsMemberLevel(uid, &pmsBrand)
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
// @Description delete the UmsMemberLevel
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UmsMemberLevelController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteUmsMemberLevel(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
