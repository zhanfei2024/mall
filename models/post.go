package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

/*
Post 结构体
*/
type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	User  *User  `orm:"rel(fk)"`  // OneToMany relation
	Tags  []*Tag `orm:"rel(m2m)"` // ManyToMany relation
}

// TableName 自定义表名
func (u *Post) TableName() string {
	return "post"
}

// AddPost 新增邮箱
func AddPost(u Post) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 插入主表
	user := User{
		Username: u.User.Username,
		Password: u.User.Password,
		Profile:  u.User.Profile,
	}

	id, err = o.Insert(&user)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 插入子表
	post := Post{
		Title: u.Title,
		User:  &User{Id: int(id)},
	}

	_, err = o.Insert(&post)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

// GetPost 查询指定邮件
func GetPost(uid int) (u *Post, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	post := &Post{Id: uid}
	err = o.Read(post)

	// 已经取得了 Post 对象，查询 User
	if post.User != nil {
		err = o.Read(post.User)
	}

	return post, err
}

// GetAllPosts 分页查询post
func GetAllPosts(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// post := new(Post)
	var posts []*Post
	qs := o.QueryTable("post")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&posts)
	for _, u := range posts {
		if u.User != nil {
			err = o.Read(u.User)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, posts), err
}

// UpdatePost 更新指定邮箱
func UpdatePost(uid int, uu *Post) (a *Post, err error) {
	o := orm.NewOrm()
	post := Post{Id: uid}
	user := User{Id: uid}

	if o.Read(&post) == nil {

		if uu.Title != "" {
			post.Title = uu.Title
		}

		if o.Read(&user) == nil {
			if uu.User.Username != "" {
				user.Username = uu.User.Username
			}
			if uu.User.Password != "" {
				user.Password = uu.User.Password
			}
			/* if uu.Profile.Age > 0 {
				profile.Age = uu.Profile.Age
			}
			if uu.Profile.Address != "" {
				profile.Address = uu.Profile.Address
			}
			if uu.Profile.Gender == 0 || uu.Profile.Gender == 1 {
				profile.Gender = uu.Profile.Gender
			}
			if uu.Profile.Email != "" {
				profile.Email = uu.Profile.Email
			} */
		}

		post.User = &user

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&post); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&user); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &post, nil
	}

	return nil, err
}

// DeletePost  删除指定邮箱
func DeletePost(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除子表
	post := Post{Id: uid}
	_, err = o.Delete(&post)
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
	// 注册 model:
	orm.RegisterModel(new(Post))
}
