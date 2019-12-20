package models

import (
	"errors"
	"mall/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// UmsMemberStatisticsInfo 结构体
type UmsMemberStatisticsInfo struct {
	Id                  int        `json:"id"`
	ConsumeAmount       float64    `orm:"digits(10);decimals(2)" description:"消费金额" json:"consume_amount"`
	OrderCount          int        `description:"订单总数" json:"order_count"`
	CouponCount         int        `description:"优惠券总数" json:"coupon_count"`
	CommentCount        int        `description:"评论总数" json:"comment_count"`
	ReturnOrderCount    int        `description:"退货订单数" json:"return_order_count"`
	LoginCount          int        `description:"登录次数" json:"login_count"`
	AttendCount         int        `description:"参加次数" json:"attend_count"`
	FansCount           int        `description:"粉丝数" json:"fans_count"`
	CollectProductCount int        `description:"收集产品计数" json:"collect_product_count"`
	CollectSubjectCount int        `description:"收集专题数" json:"collect_subject_count"`
	CollectTopicCount   int        `description:"收集主题数" json:"collect_topic_count"`
	CollectCommentCount int        `description:"收集评论数" json:"collect_comment_count"`
	InviteFriendCount   int        `description:"邀请好友计数" json:"invite_friend_count"`
	RecentOrderTime     int        `description:"最近订购时间" json:"recent_order_time"`
	UmsMember           *UmsMember `orm:"reverse(one)" json:"-"` // 设置一对一反向关系(可选)
}

// TableName 自定义表名
func (u *UmsMemberStatisticsInfo) TableName() string {
	return "ums_member_statistics_info"
}

// AddUmsMemberStatisticsInfo 新增商品分类
func AddUmsMemberStatisticsInfo(u UmsMemberStatisticsInfo) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 插入主表
	umsMemberStatisticsInfo := UmsMemberStatisticsInfo{
		ConsumeAmount:       u.ConsumeAmount,
		OrderCount:          u.OrderCount,
		CouponCount:         u.CouponCount,
		CommentCount:        u.CommentCount,
		ReturnOrderCount:    u.ReturnOrderCount,
		LoginCount:          u.LoginCount,
		AttendCount:         u.AttendCount,
		FansCount:           u.FansCount,
		CollectProductCount: u.CollectProductCount,
		CollectSubjectCount: u.CollectSubjectCount,
		CollectTopicCount:   u.CollectTopicCount,
		CollectCommentCount: u.CollectCommentCount,
		InviteFriendCount:   u.InviteFriendCount,
		RecentOrderTime:     u.ReturnOrderCount,
	}
	// 开启事务
	err = o.Begin()

	id, err = o.Insert(&umsMemberStatisticsInfo)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

/*
GetUmsMemberStatisticsInfo 查询
*/
func GetUmsMemberStatisticsInfo(uid int) (u *UmsMemberStatisticsInfo, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	umsMemberStatisticsInfo := &UmsMemberStatisticsInfo{Id: uid}

	err = o.Read(&umsMemberStatisticsInfo)

	return umsMemberStatisticsInfo, err
}

// GetAllUmsMemberStatisticsInfo 分页查询
func GetAllUmsMemberStatisticsInfo(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	var umsMemberLevel []UmsMemberStatisticsInfo
	qs := o.QueryTable("ums_member_statistics_info")
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&umsMemberLevel)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, umsMemberLevel), err
}

// UpdateUmsMemberStatisticsInfo 更新
func UpdateUmsMemberStatisticsInfo(uid int, uu *UmsMemberStatisticsInfo) (a *UmsMemberStatisticsInfo, err error) {
	// 创建一个 orm对象
	o := orm.NewOrm()
	// 实例化 UmsMemberStatisticsInfo
	umsMemberLevel := UmsMemberStatisticsInfo{Id: uid}

	if o.Read(&umsMemberLevel) == nil {

		umsMemberStatisticsInfo := UmsMemberStatisticsInfo{
			ConsumeAmount:       uu.ConsumeAmount,
			OrderCount:          uu.OrderCount,
			CouponCount:         uu.CouponCount,
			CommentCount:        uu.CommentCount,
			ReturnOrderCount:    uu.ReturnOrderCount,
			LoginCount:          uu.LoginCount,
			AttendCount:         uu.AttendCount,
			FansCount:           uu.FansCount,
			CollectProductCount: uu.CollectProductCount,
			CollectSubjectCount: uu.CollectSubjectCount,
			CollectTopicCount:   uu.CollectTopicCount,
			CollectCommentCount: uu.CollectCommentCount,
			InviteFriendCount:   uu.InviteFriendCount,
			RecentOrderTime:     uu.RecentOrderTime,
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&umsMemberStatisticsInfo); err != nil {
			return nil, errors.New("修改失败")
		}

		if _, err := o.Update(&umsMemberStatisticsInfo); err != nil {
			return nil, errors.New("修改失败")
		}
		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &umsMemberStatisticsInfo, nil
	}

	return nil, err
}

// DeleteUmsMemberStatisticsInfo 删除
func DeleteUmsMemberStatisticsInfo(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除主表
	umsMemberStatisticsInfo := UmsMemberStatisticsInfo{Id: uid}
	_, err = o.Delete(&umsMemberStatisticsInfo)
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
	orm.RegisterModel(new(UmsMemberStatisticsInfo))
}
