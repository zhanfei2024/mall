package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Comment 定义结构体
type Comment struct {
	Id      int      `json:"id"`
	Comment string   `description:"描述" json:"comment"`
	User    *User    `orm:"rel(fk)" json:"user"`
	Monitor *Monitor `orm:"rel(fk)" json:"monitor"`
}

// TableName 自定义表名
func (c *Comment) TableName() string {
	return "comments"
}

// AddComment 新增评论
func AddComment(c Comment) (id int64, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 创建 comment 对象
	comment := Comment{
		Comment: c.Comment,
	}
	// 创建 profile 对象
	profile := Profile{Id: c.User.Id}
	// 创建 user 对象
	user := User{Id: c.User.Id, Profile: c.User.Profile}
	// 创建 monitor 对象
	monitor := Monitor{Id: c.Monitor.Id}

	if o.Read(&user) == nil {
		user.Profile = &profile
		comment.User = &user
	}

	if o.Read(&monitor) == nil {
		comment.Monitor = &monitor
	}

	// 开启事务
	o.Begin()

	// 插入评论
	id, err = o.Insert((&comment))
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}
	// 提交事务
	err = o.Commit()

	return id, err
}

// UpdateComment 更新评论
func UpdateComment(uid int, uu *Comment) (a *Comment, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()

	// 创建 comment 对象
	comment := Comment{}

	// 创建 user 对象
	user := User{Id: uu.User.Id}

	// 创建 monitor 对象
	monitor := Monitor{Id: uu.Monitor.Id}

	if o.Read(&user) == nil {
		if o.Read(&monitor) == nil {
			if uu.Comment != "" {
				comment = Comment{Id: uid, Comment: uu.Comment, User: &user, Monitor: &monitor}
			}
			// 开启事务
			err = o.Begin()
			if _, err := o.Update(&comment); err != nil {
				return nil, errors.New("修改失败")
			}

			if err != nil {
				// 事务回退
				err = o.Rollback()
			} else {
				// 提交事务
				err = o.Commit()
			}
			return &comment, nil
		}
	}
	return nil, err
}

// GetComment 查询单个评论
func GetComment(uid int) (c *Comment, err error) {
	o := orm.NewOrm()
	comment := &Comment{Id: uid}
	err = o.Read(comment)

	if comment.User != nil {
		err = o.Read(comment.User)
	}

	if comment.Monitor != nil {
		err = o.Read(comment.Monitor)
	}

	return comment, err
}

// GetAllComments 分页查询评论
func GetAllComments(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	comment := new(Comment)
	var comments []Comment
	qs := o.QueryTable(comment)
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&comments)
	for _, u := range comments {
		if u.User != nil {
			err = o.Read(u.User)
		}
		if u.Monitor != nil {
			err = o.Read(u.Monitor)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, comments), err
}

// DeleteComment 删除指定评论
func DeleteComment(uid int) (b bool, err error) {
	// 创建 oremer 实例
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除表
	comment := Comment{Id: uid}
	_, err = o.Delete(&comment)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()

	return b, err
}

// 注册 model
func init() {
	orm.RegisterModel(new(Comment))
}
