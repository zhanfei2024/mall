package controllers

import (
	"encoding/json"
	"first_go/models"
	"first_go/utils"
	"fmt"
	"strconv"
)

// Operations about UmsMember
type UmsMemberController struct {
	BaseController
}

// @Title CreateUmsMember
// @Description create UmsMember
// @Param	body		body 	models.UmsMember	true		"body for UmsMember content"
// @Success 200 {int} models.UmsMember.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UmsMemberController) Post() {
	var umsMember models.UmsMember
	fmt.Println(json.Unmarshal(u.Ctx.Input.RequestBody, &umsMember), "==========++++=========")
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &umsMember)
	uid, _ := models.AddUmsMember(umsMember)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all UmsMembers
// @Success 200 {object} models.UmsMember
// @router / [get]
func (u *UmsMemberController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllUmsMember(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get UmsMember by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UmsMember
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UmsMemberController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsBrand, err := models.GetUmsMember(uid)
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
// @Description update the UmsMember
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.UmsMember	true		"body for UmsMember content"
// @Success 200 {object} models.UmsMember
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UmsMemberController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsBrand models.UmsMember
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
		uu, err := models.UpdateUmsMember(uid, &pmsBrand)
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
// @Description delete the UmsMember
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UmsMemberController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteUmsMember(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
