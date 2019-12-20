package controllers

import (
	"encoding/json"
	"first_go/models"
	"first_go/utils"
	"strconv"
)

// Operations about PmsProductCategory
type PmsProductCategoryController struct {
	BaseController
}

// @Title CreatePmsProductCategory
// @Description create PmsProductCategory
// @Param	body		body 	models.PmsProductCategory	true		"body for PmsProductCategory content"
// @Success 200 {int} models.PmsProductCategory.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PmsProductCategoryController) Post() {
	var productGategory models.PmsProductCategory
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &productGategory)
	uid, _ := models.AddPmsProductCategory(productGategory)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all PmsProductCategorys
// @Success 200 {object} models.PmsProductCategory
// @router / [get]
func (u *PmsProductCategoryController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPmsProductCategory(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get PmsProductCategory by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PmsProductCategory
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PmsProductCategoryController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		productGategory, err := models.GetPmsProductCategory(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, productGategory}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the PmsProductCategory
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.PmsProductCategory	true		"body for PmsProductCategory content"
// @Success 200 {object} models.PmsProductCategory
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PmsProductCategoryController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var productGategory models.PmsProductCategory
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &productGategory)
		uu, err := models.UpdatePmsProductCategory(uid, &productGategory)
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
// @Description delete the PmsProductCategory
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PmsProductCategoryController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePmsProductCategory(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
