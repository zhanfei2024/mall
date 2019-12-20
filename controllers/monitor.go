package controllers

import (
	"encoding/json"
	"mall/models"
	"mall/utils"
	"strconv"
)

// Operations about Monitor
type MonitorController struct {
	BaseController
}

// @Title CreateMonitor
// @Description create monitor
// @Param	body		body 	models.Monitor	true		"body for monitor content"
// @Success 200 {int} models.Monitor.Id
// @Failure 403 body is empty
// @router / [post]
func (u *MonitorController) Post() {
	var monitor models.Monitor
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &monitor)
	uid, _ := models.AddMonitor(monitor)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Monitors
// @Success 200 {object} models.Monitor
// @router / [get]
func (u *MonitorController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllMonitor(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get monitor by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Monitor
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *MonitorController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		monitor, err := models.GetMonitor(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, monitor}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the monitor
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Monitor	true		"body for monitor content"
// @Success 200 {object} models.Monitor
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *MonitorController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var monitor models.Monitor
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &monitor)
		uu, err := models.UpdateMonitor(uid, &monitor)
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
// @Description delete the monitor
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *MonitorController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteMonitor(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
