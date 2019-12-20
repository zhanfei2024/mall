package models

import (
	"errors"
	"mall/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// UmsMemberReceiveAddress 定义结构体
type UmsMemberReceiveAddress struct {
	Id            int        `json:"id"`
	Name          string     `description:"收货人名称" orm:"size(64)" json:"name"`
	Phone         string     `description:"电话" orm:"size(64)" json:"phone"`
	DefaultStatus int        `description:"默认状态 0 非默认， 1 默认" json:"default_status"`
	PostCode      string     `description:"邮政编码" orm:"size(100)" json:"post_code"`
	Province      string     `description:"省名称" orm:"size(100)" json:"province"`
	City          string     `description:"城市名" orm:"size(100)" json:"city"`
	Region        string     `description:"区(县)" orm:"size(100) json:"region"`
	DetailAddress string     `description:"详细地址" orm:"size(100)" json:"detail_address"`
	UmsMember     *UmsMember `orm:"rel(fk)" json:"ums_member"`
}

// TableName 自定义表名
func (c *UmsMemberReceiveAddress) TableName() string {
	return "ums_member_receive_address"
}

// AddUmsMemberReceiveAddress 新增
func AddUmsMemberReceiveAddress(c UmsMemberReceiveAddress) (id int64, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 基本实例化 umsMember 结构体
	umsMember := UmsMember{Id: c.UmsMember.Id}
	// &umsMember  指针类型的结构体
	umsMemberErr := o.Read(&umsMember)

	umsMemberReceiveAddress := UmsMemberReceiveAddress{
		UmsMember:     &umsMember,
		Name:          c.Name,
		Phone:         c.Phone,
		DefaultStatus: c.DefaultStatus,
		PostCode:      c.PostCode,
		Province:      c.Province,
		City:          c.City,
		Region:        c.Region,
		DetailAddress: c.DetailAddress,
	}

	if umsMemberErr == nil {
		// 开启事务
		o.Begin()

		// 插入评论
		id, err = o.Insert(&umsMemberReceiveAddress)
		if err != nil {
			// 回滚事务
			err = o.Rollback()
		}
		// 提交事务
		err = o.Commit()

		return id, err

	}
	return -1, umsMemberErr
}

// UpdateUmsMemberReceiveAddress 更新
func UpdateUmsMemberReceiveAddress(uid int, uu *UmsMemberReceiveAddress) (a *UmsMemberReceiveAddress, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 创建 umsMember 对象
	umsMember := UmsMember{Id: uu.UmsMember.Id}

	umsMemberErr := o.Read(&umsMember)

	if umsMemberErr == nil {
		umsMemberReceiveAddress := UmsMemberReceiveAddress{
			Id:            uid,
			UmsMember:     &umsMember,
			Name:          uu.Name,
			Phone:         uu.Phone,
			DefaultStatus: uu.DefaultStatus,
			PostCode:      uu.PostCode,
			Province:      uu.Province,
			City:          uu.City,
			Region:        uu.Region,
			DetailAddress: uu.DetailAddress,
		}
		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&umsMemberReceiveAddress); err != nil {
			return nil, errors.New("修改失败")
		}

		if err != nil {
			// 事务回退
			err = o.Rollback()
		} else {
			// 提交事务
			err = o.Commit()
		}
		return &umsMemberReceiveAddress, nil
	}
	return nil, err
}

// GetUmsMemberReceiveAddress 查询单个
func GetUmsMemberReceiveAddress(uid int) (c *UmsMemberReceiveAddress, err error) {
	o := orm.NewOrm()
	umsMemberReceiveAddress := UmsMemberReceiveAddress{Id: uid}

	err = o.Read(&umsMemberReceiveAddress)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		if umsMemberReceiveAddress.UmsMember != nil {
			err = o.Read(umsMemberReceiveAddress.UmsMember)
		}

	}

	return &umsMemberReceiveAddress, err

}

// GetAllUmsMemberReceiveAddress 分页查询评论
func GetAllUmsMemberReceiveAddress(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()

	var umsMemberReceiveAddress []*UmsMemberReceiveAddress

	qs := o.QueryTable("ums_member_receive_address")

	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&umsMemberReceiveAddress)
	for _, u := range umsMemberReceiveAddress {
		if u.UmsMember != nil {
			err = o.Read(u.UmsMember)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, umsMemberReceiveAddress), err
}

// DeleteUmsMemberReceiveAddress 删除指定评论
func DeleteUmsMemberReceiveAddress(uid int) (b bool, err error) {
	// 创建 oremer 实例
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除表
	pmsProduct := UmsMemberReceiveAddress{Id: uid}
	_, err = o.Delete(&pmsProduct)

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
	orm.RegisterModel(new(UmsMemberReceiveAddress))
}
