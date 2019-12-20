package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about PmsProducts
type PmsProductController struct {
	BaseController
}

// @Title CreatePmsProduct
// @Description create pmsProduct
// @Param	body		body 	models.PmsProduct	true		"body for pmsProduct content"
// @Success 200 {int} models.pmsProduct.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PmsProductController) Post() {
	var pmsProduct models.PmsProduct
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsProduct)

	uid, _ := models.AddPmsProduct(pmsProduct)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all pmsProducts
// @Success 200 {object} models.PmsProduct
// @router / [get]
func (u *PmsProductController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPmsProducts(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get pmsProduct by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.PmsProduct
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PmsProductController) Get() {

	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		pmsProduct, err := models.GetPmsProduct(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, pmsProduct}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the pmsProduct
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.PmsProduct	true		"body for pmsProduct content"
// @Success 200 {object} models.PmsProduct
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PmsProductController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var pmsProduct models.PmsProduct
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &pmsProduct)
		uu, err := models.UpdatePmsProduct(uid, &pmsProduct)
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
// @Description delete the pmsProduct
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PmsProductController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePmsProduct(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
