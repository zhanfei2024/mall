package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// PmsBrand 结构体
type PmsBrand struct {
	Id                  int    `json:"id"`
	Name                string `orm:"size(64)" description:"品牌名称" json:"name"`
	FirstLetter         string `orm:"size(8)" description:"首字母" json:"first_letter"`
	Sort                int    `description:"排序" json:"sort"`
	FactoryStatus       int    `description:"是否为品牌制造商：0->不是；1->是" json:"factory_status"`
	ShowStatus          int    `description:"是否显示" json:"show_status"`
	ProductCount        int    `description:"产品数量" json:"product_count"`
	ProductCommentCount int    `description:"产品评论数量" json:"product_comment_count"`
	Logo                string `description:"品牌logo" json:"logo"`
	BigPic              string `description:"专区大图" json:"big_pic"`
	BrandStory          string `orm:"type(text)" description:"品牌故事" json:"brand_story"`
	// PmsProduct          []*PmsProduct `orm:"reverse(many)"`
}

// TableName 自定义表名
func (u *PmsBrand) TableName() string {
	return "pms_brand"
}

// AddPmsBrand 新增商品分类
func AddPmsBrand(u PmsBrand) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()
	// 插入主表
	pmsBrand := PmsBrand{
		Name:                u.Name,
		FirstLetter:         u.FirstLetter,
		Sort:                u.Sort,
		FactoryStatus:       u.FactoryStatus,
		ShowStatus:          u.ShowStatus,
		ProductCount:        u.ProductCount,
		ProductCommentCount: u.ProductCommentCount,
		Logo:                u.Logo,
		BigPic:              u.BigPic,
		BrandStory:          u.BrandStory,
	}

	id, err = o.Insert(&pmsBrand)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

/*
GetPmsBrand 查询商品分类
*/
func GetPmsBrand(uid int) (u *PmsBrand, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	pmsBrand := &PmsBrand{Id: uid}
	err = o.Read(pmsBrand)

	return pmsBrand, err
}

// GetAllPmsBrand 分页查询商品分类
func GetAllPmsBrand(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	// user := new(User)
	var pmsBrand []PmsBrand
	qs := o.QueryTable("pms_brand")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&pmsBrand)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, pmsBrand), err
}

// UpdatePmsBrand 更新
func UpdatePmsBrand(uid int, uu *PmsBrand) (a *PmsBrand, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 PmsBrand
	pmsBrand := PmsBrand{Id: uid}

	if o.Read(&pmsBrand) == nil {

		if uu.Name != "" {
			pmsBrand.Name = uu.Name
		}
		if uu.FirstLetter != "" {
			pmsBrand.FirstLetter = uu.FirstLetter
		}
		if uu.Sort >= 0 {
			pmsBrand.Sort = uu.Sort
		}
		if uu.FactoryStatus >= 0 {
			pmsBrand.FactoryStatus = uu.FactoryStatus
		}
		if uu.ShowStatus >= 0 {
			pmsBrand.ShowStatus = uu.ShowStatus
		}
		if uu.ProductCount >= 0 {
			pmsBrand.ProductCount = uu.ProductCount
		}
		if uu.ProductCommentCount >= 0 {
			pmsBrand.ProductCommentCount = uu.ProductCommentCount
		}
		if uu.Logo != "" {
			pmsBrand.Logo = uu.Logo
		}
		if uu.BigPic != "" {
			pmsBrand.BigPic = uu.BigPic
		}
		if uu.BrandStory != "" {
			pmsBrand.BrandStory = uu.BrandStory
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&pmsBrand); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&pmsBrand); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &pmsBrand, nil
	}

	return nil, err
}

// DeletePmsBrand 删除
func DeletePmsBrand(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	pmsBrand := PmsBrand{Id: uid}
	_, err = o.Delete(&pmsBrand)
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
	orm.RegisterModel(new(PmsBrand))
}
