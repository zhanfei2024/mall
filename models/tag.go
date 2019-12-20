package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Tag 结构体
type Tag struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Posts []*Post `orm:"reverse(many)"`
}

// TableName table名称
func (u *Tag) TableName() string {
	return "tag"
}

// AddTag 新增Tag
func AddTag(u Tag) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 插入
	tag := Tag{
		Id:   u.Id,
		Name: u.Name,
	}

	id, err = o.Insert(&tag)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

// GetTag 查询tag
func GetTag(uid int) (m *Tag, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	tag := &Tag{Id: uid}
	err = o.Read(tag)

	return tag, err
}

// GetAllTag 分页查询tag
func GetAllTag(p int, size int) (u utils.Page, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()
	tag := new(Tag)
	var tags []Tag
	qs := o.QueryTable(tag)
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&tags)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, tags), err
}

// UpdateTag 更新指定tag
func UpdateTag(uid int, mm *Tag) (a *Tag, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()
	tag := Tag{Id: uid}

	if o.Read(&tag) == nil {

		if mm.Name != "" {
			tag.Name = mm.Name
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&tag); err != nil {
			return nil, errors.New("修改失败")
		}

		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &tag, nil
	}

	return nil, err
}

// DeleteTag 删除指定用户
func DeleteTag(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除
	tag := Tag{Id: uid}
	_, err = o.Delete(&tag)
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
	orm.RegisterModel(new(Tag))
}
