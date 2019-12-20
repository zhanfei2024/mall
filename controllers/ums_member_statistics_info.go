package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about UmsMemberStatisticsInfo
type UmsMemberStatisticsInfoController struct {
	BaseController
}

// @Title CreateUmsMemberStatisticsInfo
// @Description create UmsMemberStatisticsInfo
// @Param	body		body 	models.UmsMemberStatisticsInfo	true		"body for UmsMemberStatisticsInfo content"
// @Success 200 {int} models.UmsMemberStatisticsInfo.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UmsMemberStatisticsInfoController) Post() {
	var pmsBrand models.UmsMemberStatisticsInfo
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
	uid, _ := models.AddUmsMemberStatisticsInfo(pmsBrand)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all UmsMemberStatisticsInfos
// @Success 200 {object} models.UmsMemberStatisticsInfo
// @router / [get]
func (u *UmsMemberStatisticsInfoController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllUmsMemberStatisticsInfo(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get UmsMemberStatisticsInfo by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UmsMemberStatisticsInfo
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UmsMemberStatisticsInfoController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsBrand, err := models.GetUmsMemberStatisticsInfo(uid)
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
// @Description update the UmsMemberStatisticsInfo
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.UmsMemberStatisticsInfo	true		"body for UmsMemberStatisticsInfo content"
// @Success 200 {object} models.UmsMemberStatisticsInfo
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UmsMemberStatisticsInfoController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsBrand models.UmsMemberStatisticsInfo
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
		uu, err := models.UpdateUmsMemberStatisticsInfo(uid, &pmsBrand)
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
// @Description delete the UmsMemberStatisticsInfo
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UmsMemberStatisticsInfoController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteUmsMemberStatisticsInfo(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
