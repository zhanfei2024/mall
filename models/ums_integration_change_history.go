package models

import (
	"errors"
	"first_go/utils"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// UmsIntegrationChangeHistory 结构体
type UmsIntegrationChangeHistory struct {
	Id          int        `json:"id"`
	CreateTime  time.Time  `orm:type(datetime) description:"修改时间" json:"create_time"`
	ChangeType  int        `description:"变更类型" json:"change_type"`
	ChangeCount int        `description:"变更次数" json:"change_count"`
	OperateMan  string     `orm:"size(100) "description:"操作人" json:"operate_man"`
	OperateNote string     `orm:"size(100) "description:"操作备注" json:"operate_note"`
	SourceType  int        `description:"操作来源" json:"source_type"`
	UmsMember   *UmsMember `orm:"rel(fk)"`
}

// TableName 自定义表名
func (u *UmsIntegrationChangeHistory) TableName() string {
	return "ums_integration_change_history"
}

// AddUmsIntegrationChangeHistory 新增
func AddUmsIntegrationChangeHistory(u UmsIntegrationChangeHistory) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsMember := UmsMember{
		Id: u.UmsMember.Id,
	}

	umsMemberErr := o.Read(&umsMember)

	if umsMemberErr == nil {
		umsGrowthChangeHistory := UmsIntegrationChangeHistory{
			UmsMember:   &umsMember,
			CreateTime:  time.Now(),
			ChangeType:  u.ChangeType,
			ChangeCount: u.ChangeCount,
			OperateMan:  u.OperateMan,
			OperateNote: u.OperateNote,
			SourceType:  u.SourceType,
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
GetUmsIntegrationChangeHistory 查询
*/
func GetUmsIntegrationChangeHistory(uid int) (u *UmsIntegrationChangeHistory, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsGrowthChangeHistory := &UmsIntegrationChangeHistory{Id: uid}

	err = o.Read(umsGrowthChangeHistory)

	if umsGrowthChangeHistory.UmsMember != nil {
		err = o.Read(umsGrowthChangeHistory.UmsMember)
	}

	return umsGrowthChangeHistory, err
}

// GetAllUmsIntegrationChangeHistory 分页查询
func GetAllUmsIntegrationChangeHistory(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	var umsGrowthChangeHistory []UmsIntegrationChangeHistory
	qs := o.QueryTable("ums_growth_change_history")
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

// UpdateUmsIntegrationChangeHistory 更新
func UpdateUmsIntegrationChangeHistory(uid int, uu *UmsIntegrationChangeHistory) (a *UmsIntegrationChangeHistory, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化
	umsMember := UmsMember{Id: uu.UmsMember.Id}

	if o.Read(&umsMember) == nil {

		umsGrowthChangeHistory := UmsIntegrationChangeHistory{
			UmsMember:   &umsMember,
			CreateTime:  time.Now(),
			ChangeType:  uu.ChangeType,
			ChangeCount: uu.ChangeCount,
			OperateMan:  uu.OperateMan,
			OperateNote: uu.OperateNote,
			SourceType:  uu.SourceType,
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

// DeleteUmsIntegrationChangeHistory 删除
func DeleteUmsIntegrationChangeHistory(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	umsGrowthChangeHistory := UmsIntegrationChangeHistory{Id: uid}
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
	orm.RegisterModel(new(UmsIntegrationChangeHistory))
}
