package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about PmsBrand
type PmsBrandController struct {
	BaseController
}

// @Title CreatePmsBrand
// @Description create PmsBrand
// @Param	body		body 	models.PmsBrand	true		"body for PmsBrand content"
// @Success 200 {int} models.PmsBrand.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PmsBrandController) Post() {
	var pmsBrand models.PmsBrand
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
	uid, _ := models.AddPmsBrand(pmsBrand)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all PmsBrands
// @Success 200 {object} models.PmsBrand
// @router / [get]
func (u *PmsBrandController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPmsBrand(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get PmsBrand by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PmsBrand
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PmsBrandController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsBrand, err := models.GetPmsBrand(uid)
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
// @Description update the PmsBrand
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.PmsBrand	true		"body for PmsBrand content"
// @Success 200 {object} models.PmsBrand
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PmsBrandController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsBrand models.PmsBrand
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsBrand)
		uu, err := models.UpdatePmsBrand(uid, &pmsBrand)
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
// @Description delete the PmsBrand
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PmsBrandController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePmsBrand(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
