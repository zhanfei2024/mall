package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about PmsProductAttributeCategory
type PmsProductAttributeCategoryController struct {
	BaseController
}

// @Title CreatePmsProductAttributeCategory
// @Description create PmsProductAttributeCategory
// @Param	body		body 	models.PmsProductAttributeCategory	true		"body for PmsProductAttributeCategory content"
// @Success 200 {int} models.PmsProductAttributeCategory.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PmsProductAttributeCategoryController) Post() {
	var pmsProductAttributeCategory models.PmsProductAttributeCategory
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsProductAttributeCategory)
	uid, _ := models.AddPmsProductAttributeCategory(pmsProductAttributeCategory)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all PmsProductAttributeCategorys
// @Success 200 {object} models.PmsProductAttributeCategory
// @router / [get]
func (u *PmsProductAttributeCategoryController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPmsProductAttributeCategory(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get PmsProductAttributeCategory by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PmsProductAttributeCategory
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PmsProductAttributeCategoryController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsProductAttributeCategory, err := models.GetPmsProductAttributeCategory(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, pmsProductAttributeCategory}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the PmsProductAttributeCategory
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.PmsProductAttributeCategory	true		"body for PmsProductAttributeCategory content"
// @Success 200 {object} models.PmsProductAttributeCategory
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PmsProductAttributeCategoryController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsProductAttributeCategory models.PmsProductAttributeCategory
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsProductAttributeCategory)
		uu, err := models.UpdatePmsProductAttributeCategory(uid, &pmsProductAttributeCategory)
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
// @Description delete the PmsProductAttributeCategory
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PmsProductAttributeCategoryController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePmsProductAttributeCategory(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
