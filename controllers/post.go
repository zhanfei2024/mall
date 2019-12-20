package controllers

import (
	"encoding/json"
	"first_go/models"
	"first_go/utils"
	"strconv"
)

// Operations about Posts
type PostController struct {
	BaseController
}

// @Title CreatePost
// @Description create post
// @Param	body		body 	models.Post	true		"body for post content"
// @Success 200 {int} models.Post.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PostController) Post() {
	var post models.Post
	_ = json.Unmarshal(u.Ctx.Input.RequestBody, &post)
	uid, _ := models.AddPost(post)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Posts
// @Success 200 {object} models.Post
// @router / [get]
func (u *PostController) GetAll() {
	currentPage, _ := strconv.Atoi(u.Ctx.Input.Query("page"))
	if currentPage == 0 {
		currentPage = 1
	}
	pageSize := utils.PageSize
	d, err := models.GetAllPosts(currentPage, pageSize)
	code, message := DecodeErr(err)

	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, d}
	}
	u.ServeJSON()
}

// @Title Get
// @Description get post by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Post
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *PostController) Get() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		user, err := models.GetUser(uid)
		code, message := DecodeErr(err)
		if err != nil {
			u.Data["json"] = ErrResponse{code, message}
		} else {
			u.Data["json"] = Response{code, message, user}
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the post
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.Post	true		"body for post content"
// @Success 200 {object} models.Post
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *PostController) Put() {
	uid, _ := u.GetInt(":uid")
	if uid > 0 {
		var post models.Post
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &post)
		uu, err := models.UpdatePost(uid, &post)
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
// @Description delete the post
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *PostController) Delete() {
	uid, _ := u.GetInt(":uid")
	b, err := models.DeletePost(uid)
	code, message := DecodeErr(err)
	if err != nil {
		u.Data["json"] = ErrResponse{code, message}
	} else {
		u.Data["json"] = Response{code, message, b}
	}
	u.ServeJSON()
}
