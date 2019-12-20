package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about UmsMemberReceiveAddress
type UmsMemberReceiveAddressController struct {
	BaseController
}

// @Title CreateUmsMemberReceiveAddress
// @Description create UmsMemberReceiveAddress
// @Param	body		body 	models.UmsMemberReceiveAddress	true		"body for UmsMemberReceiveAddress content"
// @Success 200 {int} models.UmsMemberReceiveAddress.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UmsMemberReceiveAddressController) Post() {
	var pmsBrand models.UmsMemberReceiveAddress
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
	uid, _ := models.AddUmsMemberReceiveAddress(pmsBrand)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all UmsMemberReceiveAddresss
// @Success 200 {object} models.UmsMemberReceiveAddress
// @router / [get]
func (u *UmsMemberReceiveAddressController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllUmsMemberReceiveAddress(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get UmsMemberReceiveAddress by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UmsMemberReceiveAddress
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UmsMemberReceiveAddressController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsBrand, err := models.GetUmsMemberReceiveAddress(uid)
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
// @Description update the UmsMemberReceiveAddress
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.UmsMemberReceiveAddress	true		"body for UmsMemberReceiveAddress content"
// @Success 200 {object} models.UmsMemberReceiveAddress
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UmsMemberReceiveAddressController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsBrand models.UmsMemberReceiveAddress
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
		uu, err := models.UpdateUmsMemberReceiveAddress(uid, &pmsBrand)
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
// @Description delete the UmsMemberReceiveAddress
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UmsMemberReceiveAddressController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteUmsMemberReceiveAddress(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
