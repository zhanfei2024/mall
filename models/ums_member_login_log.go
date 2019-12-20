package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// UmsMemberLoginLog 结构体
type UmsMemberLoginLog struct {
	Id         int        `json:"id"`
	CreateTime int        `description:"修改时间" json:"create_time"`
	Ip         string     `orm:"size(64)" description:"变更类型" json:"ip"`
	City       string     `orm:"size(64)" description:"城市" json:"city"`
	Province   string     `orm:"size(64)" description:"省" json:"province"`
	LoginType  int        `json:"login_type"`
	UmsMember  *UmsMember `orm:"rel(fk)" json:"ums_member"`
}

// TableName 自定义表名
func (u *UmsMemberLoginLog) TableName() string {
	return "ums_member_login_log"
}

// AddUmsMemberLoginLog 新增
func AddUmsMemberLoginLog(u UmsMemberLoginLog) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsMember := UmsMember{
		Id: u.UmsMember.Id,
	}

	umsMemberErr := o.Read(&umsMember)

	if umsMemberErr == nil {
		umsGrowthChangeHistory := UmsMemberLoginLog{
			UmsMember:  &umsMember,
			CreateTime: u.CreateTime,
			Ip:         u.Ip,
			City:       u.City,
			Province:   u.Province,
			LoginType:  u.LoginType,
		}
		// 开启事务
		err = o.Begin()

		id, err = o.Insert(&umsGrowthChangeHistory)

		if err != nil {
			// 回滚事务
			err = o.Rollback()
		}

		// 提交事务
		err = o.Commit()
		return id, err

	}
	return 0, err
}

/*
GetUmsMemberLoginLog 查询
*/
func GetUmsMemberLoginLog(uid int) (u *UmsMemberLoginLog, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsGrowthChangeHistory := &UmsMemberLoginLog{Id: uid}

	err = o.Read(umsGrowthChangeHistory)

	if umsGrowthChangeHistory.UmsMember != nil {
		err = o.Read(umsGrowthChangeHistory.UmsMember)
	}

	return umsGrowthChangeHistory, err
}

// GetAllUmsMemberLoginLog 分页查询
func GetAllUmsMemberLoginLog(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	var umsGrowthChangeHistory []UmsMemberLoginLog
	qs := o.QueryTable("ums_member_login_log")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&umsGrowthChangeHistory)
	for _, u := range umsGrowthChangeHistory {
		if u.UmsMember != nil {
			err = o.Read(u.UmsMember)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, umsGrowthChangeHistory), err
}

// UpdateUmsMemberLoginLog 更新
func UpdateUmsMemberLoginLog(uid int, uu *UmsMemberLoginLog) (a *UmsMemberLoginLog, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化
	umsMember := UmsMember{Id: uu.UmsMember.Id}

	if o.Read(&umsMember) == nil {

		umsGrowthChangeHistory := UmsMemberLoginLog{
			UmsMember:  &umsMember,
			CreateTime: uu.CreateTime,
			Ip:         uu.Ip,
			City:       uu.City,
			Province:   uu.Province,
			LoginType:  uu.LoginType,
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&umsGrowthChangeHistory); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&umsGrowthChangeHistory); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &umsGrowthChangeHistory, nil
	}

	return nil, err
}

// DeleteUmsMemberLoginLog 删除
func DeleteUmsMemberLoginLog(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	umsGrowthChangeHistory := UmsMemberLoginLog{Id: uid}
	_, err = o.Delete(&umsGrowthChangeHistory)
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
	orm.RegisterModel(new(UmsMemberLoginLog))
}
