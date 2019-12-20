package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// UmsMemberLevel 结构体
type UmsMemberLevel struct {
	Id                    int        `json:"id"`
	Name                  string     `orm:"size(100)" description:"级别名称" json:"name"`
	GrowthPoint           int        `description:"成长值" json:"growth_point"`
	DefaultStatus         int        `description:"默认状态" json:"default_status"`
	FreeFreightPoint      int        `description:"可用免费快递点数" json:"free_freight_point"`
	CommentGrowthPoint    int        `description:"点评获得的点数" json:"comment_growth_point"`
	PriviledgeFreeFreight int        `description:"减免运费特权 1 获取， 0 未获取" json:"priviledge_free_freight"`
	PriviledgeSifnIn      int        `description:"签到奖励 1 获取， 0 未获取" json:"priviledge_sifn_in"`
	PriviledgeComment     int        `description:"评价奖励 1 获取， 0 未获取" json:"priviledge_comment"`
	PriviledgePromotion   int        `description:"专享活动 1 获取， 0 未获取" json:"priviledge_promotion"`
	PriviledgeMemberPrice int        `description:"会员特价 1 获取， 0 未获取" json:"priviledge_member_price"`
	PriviledgeBirthday    int        `description:"生日礼包 1 获取， 0 未获取" json:"priviledge_birthday"`
	Note                  string     `orm:"size(200)" description:"备注" json:"note"`
	UmsMember             *UmsMember `orm:"reverse(one)" json:"-"` // 设置一对一反向关系(可选)
}

// TableName 自定义表名
func (u *UmsMemberLevel) TableName() string {
	return "ums_member_level"
}

// AddUmsMemberLevel 新增商品分类
func AddUmsMemberLevel(u UmsMemberLevel) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 插入主表
	umsMemberlevel := UmsMemberLevel{
		Name:                  u.Name,
		GrowthPoint:           u.GrowthPoint,
		DefaultStatus:         u.DefaultStatus,
		FreeFreightPoint:      u.FreeFreightPoint,
		CommentGrowthPoint:    u.CommentGrowthPoint,
		PriviledgeFreeFreight: u.PriviledgeFreeFreight,
		PriviledgeSifnIn:      u.PriviledgeSifnIn,
		PriviledgeComment:     u.PriviledgeComment,
		PriviledgePromotion:   u.PriviledgePromotion,
		PriviledgeMemberPrice: u.PriviledgeMemberPrice,
		PriviledgeBirthday:    u.PriviledgeBirthday,
		Note:                  u.Note,
	}
	// 开启事务
	err = o.Begin()

	id, err = o.Insert(&umsMemberlevel)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

/*
GetUmsMemberLevel 查询
*/
func GetUmsMemberLevel(uid int) (u *UmsMemberLevel, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsMemberLevel := &UmsMemberLevel{Id: uid}

	err = o.Read(&umsMemberLevel)

	return umsMemberLevel, err
}

// GetAllUmsMemberLevel 分页查询
func GetAllUmsMemberLevel(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	var umsMemberLevel []UmsMemberLevel
	qs := o.QueryTable("ums_member_level")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&umsMemberLevel)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, umsMemberLevel), err
}

// UpdateUmsMemberLevel 更新
func UpdateUmsMemberLevel(uid int, uu *UmsMemberLevel) (a *UmsMemberLevel, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 UmsMemberLevel
	umsMemberLevel := UmsMemberLevel{Id: uid}

	if o.Read(&umsMemberLevel) == nil {

		umsMemberlevel := UmsMemberLevel{
			Name:                  uu.Name,
			GrowthPoint:           uu.GrowthPoint,
			DefaultStatus:         uu.DefaultStatus,
			FreeFreightPoint:      uu.FreeFreightPoint,
			CommentGrowthPoint:    uu.CommentGrowthPoint,
			PriviledgeFreeFreight: uu.PriviledgeFreeFreight,
			PriviledgeSifnIn:      uu.PriviledgeSifnIn,
			PriviledgeComment:     uu.PriviledgeComment,
			PriviledgePromotion:   uu.PriviledgePromotion,
			PriviledgeMemberPrice: uu.PriviledgeMemberPrice,
			PriviledgeBirthday:    uu.PriviledgeBirthday,
			Note:                  uu.Note,
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&umsMemberlevel); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&umsMemberlevel); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &umsMemberlevel, nil
	}

	return nil, err
}

// DeleteUmsMemberLevel 删除
func DeleteUmsMemberLevel(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	umsMemberlevel := UmsMemberLevel{Id: uid}
	_, err = o.Delete(&umsMemberlevel)
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
	orm.RegisterModel(new(UmsMemberLevel))
}
