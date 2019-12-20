package models

import (
	"errors"
	"first_go/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// PmsFeightTemplate 结构体
type PmsFeightTemplate struct {
	Id             int     `json:"id"`
	Name           string  `description:"运费模板名称" json:"name"`
	ChargeType     int     `description:"收费类型" json:"charge_type"`
	FirstWeight    float64 `description:"首重" orm:"digits(10);decimals(2)" json:"first_weight"`
	FirstFee       float64 `description:"首费" orm:"digits(10);decimals(2)" json:"first_fee"`
	ContinueWeight float64 `description:"续重" orm:"digits(10);decimals(2)" json:"continue_weight"`
	ContinueFee    float64 `description:"续费" orm:"digits(10);decimals(2)" json:"continue_fee"`
	Dest           string  `description:"目的地" json:"dest"`
	// PmsProduct     []*PmsProduct `orm:"reverse(many)"`
}

// TableName 自定义表名
func (u *PmsFeightTemplate) TableName() string {
	return "pms_feight_template"
}

// AddPmsFeightTemplate 新增用户
func AddPmsFeightTemplate(u PmsFeightTemplate) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()
	// 插入主表
	pmsFeightTemplate := PmsFeightTemplate{
		Name:           u.Name,
		ChargeType:     u.ChargeType,
		FirstWeight:    u.FirstWeight,
		FirstFee:       u.FirstFee,
		ContinueWeight: u.ContinueWeight,
		ContinueFee:    u.ContinueFee,
		Dest:           u.Dest,
	}

	id, err = o.Insert(&pmsFeightTemplate)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

// GetPmsFeightTemplate 查询指定用户
func GetPmsFeightTemplate(uid int) (u *PmsFeightTemplate, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	pmsFeightTemplate := &PmsFeightTemplate{Id: uid}
	err = o.Read(pmsFeightTemplate)

	return pmsFeightTemplate, err
}

// GetAllPmsFeightTemplate 分页查询用户
func GetAllPmsFeightTemplate(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// user := new(PmsFeightTemplate)
	var pmsFeightTemplates []PmsFeightTemplate
	qs := o.QueryTable("pms_feight_template")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&pmsFeightTemplates)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, pmsFeightTemplates), err
}

// UpdatePmsFeightTemplate 更新
func UpdatePmsFeightTemplate(uid int, uu *PmsFeightTemplate) (a *PmsFeightTemplate, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化
	pmsFeightTemplates := PmsFeightTemplate{Id: uid}

	if o.Read(&pmsFeightTemplates) == nil {

		if uu.Name != "" {
			pmsFeightTemplates.Name = uu.Name
		}
		if uu.ChargeType >= 0 {
			pmsFeightTemplates.ChargeType = uu.ChargeType
		}
		if uu.FirstWeight >= 0 {
			pmsFeightTemplates.FirstWeight = uu.FirstWeight
		}
		if uu.FirstFee >= 0 {
			pmsFeightTemplates.FirstFee = uu.FirstFee
		}
		if uu.ContinueWeight >= 0 {
			pmsFeightTemplates.ContinueWeight = uu.ContinueWeight
		}
		if uu.ContinueFee >= 0 {
			pmsFeightTemplates.ContinueFee = uu.ContinueFee
		}
		if uu.Dest != "" {
			pmsFeightTemplates.Dest = uu.Dest
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&pmsFeightTemplates); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&pmsFeightTemplates); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &pmsFeightTemplates, nil
	}

	return nil, err
}

// DeletePmsFeightTemplate 删除指定用户
func DeletePmsFeightTemplate(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	pmsFeightTemplates := PmsFeightTemplate{Id: uid}
	_, err = o.Delete(&pmsFeightTemplates)
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
	orm.RegisterModel(new(PmsFeightTemplate))
}
