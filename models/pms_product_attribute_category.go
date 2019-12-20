package models

import (
	"errors"
	"first_go/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// PmsProductAttributeCategory 结构体
type PmsProductAttributeCategory struct {
	Id             int    `json:"id"`
	Name           string `orm:"size(64)" json:"name"`
	AttributeCount int    `description:"属性数量" json:"attribute_count"`
	ParamCount     int    `description:"参数数量" json:"param_count"`
	// PmsProduct     []*PmsProduct `orm:"reverse(many)"`
}

// TableName 自定义表名
func (u *PmsProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

// AddPmsProductAttributeCategory 新增商品分类
func AddPmsProductAttributeCategory(u PmsProductAttributeCategory) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()
	// 插入主表
	pmsProductAttributeCategory := PmsProductAttributeCategory{
		Name:           u.Name,
		AttributeCount: u.AttributeCount,
		ParamCount:     u.ParamCount,
	}

	id, err = o.Insert(&pmsProductAttributeCategory)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

/*
GetPmsProductAttributeCategory 查询商品分类
*/
func GetPmsProductAttributeCategory(uid int) (u *PmsProductAttributeCategory, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	pmsProductAttributeCategory := &PmsProductAttributeCategory{Id: uid}
	err = o.Read(pmsProductAttributeCategory)

	return pmsProductAttributeCategory, err
}

// GetAllPmsProductAttributeCategory 分页查询商品分类
func GetAllPmsProductAttributeCategory(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// user := new(User)
	var pmsProductAttributeCategory []PmsProductAttributeCategory
	qs := o.QueryTable("pms_brand")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&pmsProductAttributeCategory)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, pmsProductAttributeCategory), err
}

// UpdatePmsProductAttributeCategory 更新
func UpdatePmsProductAttributeCategory(uid int, uu *PmsProductAttributeCategory) (a *PmsProductAttributeCategory, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 PmsProductAttributeCategory
	pmsProductAttributeCategory := PmsProductAttributeCategory{Id: uid}

	if o.Read(&pmsProductAttributeCategory) == nil {

		if uu.Name != "" {
			pmsProductAttributeCategory.Name = uu.Name
		}
		if uu.AttributeCount >= 0 {
			pmsProductAttributeCategory.AttributeCount = uu.AttributeCount
		}
		if uu.ParamCount >= 0 {
			pmsProductAttributeCategory.ParamCount = uu.ParamCount
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&pmsProductAttributeCategory); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&pmsProductAttributeCategory); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &pmsProductAttributeCategory, nil
	}

	return nil, err
}

// DeletePmsProductAttributeCategory 删除
func DeletePmsProductAttributeCategory(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	pmsProductAttributeCategory := PmsProductAttributeCategory{Id: uid}
	_, err = o.Delete(&pmsProductAttributeCategory)
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
	orm.RegisterModel(new(PmsProductAttributeCategory))
}
