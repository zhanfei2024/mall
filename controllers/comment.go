package controllers

import (
	"encoding/json"
	"first_go/models"
	"first_go/utils"
	"strconv"
)

// Operations about Comments
type CommentController struct {
	BaseController
}

// @Title CreateComment
// @Description create comment
// @Param	body		body 	models.Comment	true		"body for comment content"
// @Success 200 {int} models.comment.Id
// @Failure 403 body is empty
// @router / [post]
func (u *CommentController) Post() {
	var comment models.Comment
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &comment)
	uid, _ := models.AddComment(comment)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all comments
// @Success 200 {object} models.Comment
// @router / [get]
func (u *CommentController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllComments(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get comment by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Comment
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *CommentController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		comment, err := models.GetComment(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, comment}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the comment
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Comment	true		"body for comment content"
// @Success 200 {object} models.Comment
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *CommentController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var comment models.Comment
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &comment)
		uu, err := models.UpdateComment(uid, &comment)
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
// @Description delete the comment
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *CommentController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeleteComment(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
