package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// PmsProductCategory 商品分类结构体
type PmsProductCategory struct {
	Id           int    `json:"id"`
	ParentId     int    `description:"上级分类的编号：0表示一级分类" json:"parent_id"`
	Name         string `description:"分类名称" orm:"size(64)" json:"name"`
	Level        int    `description:"分类级别：0->1级；1->2级" json:"level"`
	ProductCount int    `description:"商品数量" json:"product_count"`
	ProductUnit  string `description:"商品单位" orm:"size(64)" json:"product_unit"`
	NavStatus    int    `description:"是否显示在导航栏：0->不显示；1->显示" json:"nav_status"`
	ShowStatus   int    `description:"显示状态：0->不显示；1->显示" json:"show_status"`
	Sort         int    `description:"排序" json:"sort"`
	Icon         string `description:"图标" json:"icon"`
	Keywords     string `description:"关键字" json:"keywords"`
	Description  string `description:"描述" orm:"type(text)" json:"description"`
	// PmsProduct   []*PmsProduct `orm:"reverse(many)"`
}

// TableName 自定义表名
func (u *PmsProductCategory) TableName() string {
	return "pms_product_category"
}

// AddPmsProductCategory 新增商品分类
func AddPmsProductCategory(u PmsProductCategory) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()
	// 插入主表
	productCategory := PmsProductCategory{
		ParentId:     u.ParentId,
		Name:         u.Name,
		Level:        u.Level,
		ProductCount: u.ProductCount,
		ProductUnit:  u.ProductUnit,
		NavStatus:    u.NavStatus,
		ShowStatus:   u.ShowStatus,
		Sort:         u.Sort,
		Icon:         u.Icon,
		Keywords:     u.Keywords,
		Description:  u.Description,
	}

	id, err = o.Insert(&productCategory)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

/*
GetPmsProductCategory 查询商品分类
*/
func GetPmsProductCategory(uid int) (u *PmsProductCategory, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	productCategory := &PmsProductCategory{Id: uid}
	err = o.Read(productCategory)

	return productCategory, err
}

// GetAllPmsProductCategory 分页查询商品分类
func GetAllPmsProductCategory(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// user := new(User)
	var productCategories []PmsProductCategory
	qs := o.QueryTable("pms_product_category")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&productCategories)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, productCategories), err
}

// UpdatePmsProductCategory 更新
func UpdatePmsProductCategory(uid int, uu *PmsProductCategory) (a *PmsProductCategory, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 productCategory
	productCategory := PmsProductCategory{Id: uid}

	if o.Read(&productCategory) == nil {

		if uu.Name != "" {
			productCategory.Name = uu.Name
		}
		if uu.Level >= 0 {
			productCategory.Level = uu.Level
		}
		if uu.ProductCount >= 0 {
			productCategory.ProductCount = uu.ProductCount
		}
		if uu.ProductUnit != "" {
			productCategory.ProductUnit = uu.ProductUnit
		}
		if uu.Icon != "" {
			productCategory.Icon = uu.Icon
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&productCategory); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&productCategory); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &productCategory, nil
	}

	return nil, err
}

// DeletePmsProductCategory 删除
func DeletePmsProductCategory(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	productCategory := PmsProductCategory{Id: uid}
	_, err = o.Delete(&productCategory)
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
	orm.RegisterModel(new(PmsProductCategory))
}
