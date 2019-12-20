package models

import (
	"errors"
	"first_go/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// User 结构体
type User struct {
	Id       int      `json:"id"`
	Username string   `orm:"size(100)" json:"username"`
	Password string   `json:"password"`
	Profile  *Profile `orm:"rel(one)" json:"profile"` // OneToOne relation
}

// Profile 结构体
type Profile struct {
	Id      int    `json:"id"`
	Gender  int    `json:"gender"`
	Age     int    `json:"age"`
	Address string `json:"address"`
	Email   string `json:"email"`
	User    *User  `orm:"reverse(one)" json:"-"` // 设置一对一反向关系(可选)
}

// TableName 自定义表名
func (u *User) TableName() string {
	return "users"
}

// TableName 自定义表名
func (u *Profile) TableName() string {
	return "users_profiles"
}

// AddUser 新增用户
func AddUser(u User) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	fmt.Println(u.Profile)
	// 开启事务
	err = o.Begin()
	// 插入主表
	profile := Profile{
		Gender:  u.Profile.Gender,
		Age:     u.Profile.Age,
		Address: u.Profile.Address,
		Email:   u.Profile.Email,
	}

	id, err = o.Insert(&profile)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 插入子表
	user := User{
		Username: u.Username,
		Password: u.Password,
		Profile:  &Profile{Id: int(id)},
	}

	_, err = o.Insert(&user)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

// GetUser 查询指定用户
func GetUser(uid int) (u *User, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	user := &User{Id: uid}
	err = o.Read(user)

	// 已经取得了 Users 对象，查询 UserProfiles
	if user.Profile != nil {
		err = o.Read(user.Profile)
	}

	return user, err
}

// GetAllUsers 分页查询用户
func GetAllUsers(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// user := new(User)
	var users []User
	qs := o.QueryTable("users")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&users)
	for _, u := range users {
		if u.Profile != nil {
			err = o.Read(u.Profile)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, users), err
}

// UpdateUser 更新指定用户
func UpdateUser(uid int, uu *User) (a *User, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 user
	user := User{Id: uid}
	// 实例化 profile
	profile := Profile{Id: uid}

	if o.Read(&user) == nil {

		if uu.Username != "" {
			user.Username = uu.Username
		}
		if uu.Password != "" {
			user.Password = uu.Password
		}

		if o.Read(&profile) == nil {

			if uu.Profile.Age > 0 {
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
			}
		}

		user.Profile = &profile

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&user); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&profile); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &user, nil
	}

	return nil, err
}

// DeleteUser 删除指定用户
func DeleteUser(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	profile := Profile{Id: uid}
	_, err = o.Delete(&profile)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 删除子表
	user := User{Id: uid}
	_, err = o.Delete(&user)
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
	orm.RegisterModel(new(User), new(Profile))
}
