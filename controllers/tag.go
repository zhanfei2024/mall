package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about Tag
type TagController struct {
	BaseController
}

// @Title CreateTag
// @Description create tag
// @Param	body		body 	models.Tag	true		"body for tag content"
// @Success 200 {int} models.Tag.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TagController) Post() {
	var tag models.Tag
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &tag)
	uid, _ := models.AddTag(tag)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all tags
// @Success 200 {object} models.Tag
// @router / [get]
func (u *TagController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllTag(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get tag by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tag
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *TagController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		tag, err := models.GetTag(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, tag}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the tag
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Tag	true		"body for tag content"
// @Success 200 {object} models.Tag
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *TagController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var tag models.Tag
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &tag)
		uu, err := models.UpdateTag(uid, &tag)
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
// @Description delete the tag
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *TagController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteTag(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
