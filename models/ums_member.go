package models

import (
	"errors"
	"first_go/utils"
	"fmt"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// UmsMember 定义结构体
type UmsMember struct {
	Id                      int                      `json:"id"`
	UserName                string                   `description:"用户名" orm:"size(64)" json:"user_name"`
	Password                string                   `description:"密码" orm:"size(64)" json:"password"`
	NickName                string                   `description:"昵称" orm:"size(64)" json:"nick_name"`
	Phone                   string                   `description:"手机" orm:"size(64)" json:"phone"`
	Status                  int                      `description:"用户状态 0 禁用 1 可用" json:"status"`
	CreateTime              int                      `description:"创建日期" json:"create_time"`
	Icon                    string                   `description:"头像" orm:"size(500)" json:"icon"`
	Gender                  int                      `description:"性别" json:"gender"`
	Birthday                int                      `description:"生日" json:"birthday"`
	City                    string                   `description:"城市" orm:"size(64)" json:"city"`
	Job                     string                   `description:"职业" orm:"size(100)" json:"job"`
	PersonalizedSignature   string                   `description:"个性签名" orm:"size(200)" json:"personalized_signature"`
	SourceType              int                      `description:"来源渠道 1 app, 2 小程序" json:"source_type"`
	Integration             int                      `description:"整合" json:"integration"`
	Growth                  int                      `description:"成长值" json:"growth"`
	LuckeyCount             int                      `description:"幸运此次" json:"luckey_count"`
	HistoryIntegration      int                      `description:"整合次数" json:"history_integration"`
	UmsMemberLevel          *UmsMemberLevel          `orm:"rel(one)" json:"ums_member_level"`
	UmsMemberStatisticsInfo *UmsMemberStatisticsInfo `orm:"rel(one)" json:"ums_member_statistics_info"`
}

// TableName 自定义表名
func (c *UmsMember) TableName() string {
	return "ums_member"
}

// RegisterUmsMember 注册会员
func RegisterUmsMember(c UmsMember) (id int64, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 开启事务
	o.Begin()
	// 插入主表 UmsMemberLevel
	umsMemberLevel := UmsMemberLevel{
		Name:                  c.UmsMemberLevel.Name,
		GrowthPoint:           c.UmsMemberLevel.GrowthPoint,
		DefaultStatus:         c.UmsMemberLevel.DefaultStatus,
		FreeFreightPoint:      c.UmsMemberLevel.FreeFreightPoint,
		CommentGrowthPoint:    c.UmsMemberLevel.CommentGrowthPoint,
		PriviledgeFreeFreight: c.UmsMemberLevel.PriviledgeFreeFreight,
		PriviledgeSifnIn:      c.UmsMemberLevel.PriviledgeSifnIn,
		PriviledgeComment:     c.UmsMemberLevel.PriviledgeComment,
		PriviledgePromotion:   c.UmsMemberLevel.PriviledgePromotion,
		PriviledgeMemberPrice: c.UmsMemberLevel.PriviledgeMemberPrice,
		PriviledgeBirthday:    c.UmsMemberLevel.PriviledgeBirthday,
		Note:                  c.UmsMemberLevel.Note,
	}
	umsMemberLevelId, err := o.Insert(&umsMemberLevel)
	if err != nil {
		// 回滚事务
		err = o.Begin()
	}

	// 插入主表 UmsMemberStatisticsInfo
	umsMemberStatisticsInfo := UmsMemberStatisticsInfo{
		ConsumeAmount:       c.UmsMemberStatisticsInfo.ConsumeAmount,
		OrderCount:          c.UmsMemberStatisticsInfo.OrderCount,
		CouponCount:         c.UmsMemberStatisticsInfo.CouponCount,
		CommentCount:        c.UmsMemberStatisticsInfo.CommentCount,
		ReturnOrderCount:    c.UmsMemberStatisticsInfo.ReturnOrderCount,
		LoginCount:          c.UmsMemberStatisticsInfo.LoginCount,
		AttendCount:         c.UmsMemberStatisticsInfo.AttendCount,
		FansCount:           c.UmsMemberStatisticsInfo.FansCount,
		CollectProductCount: c.UmsMemberStatisticsInfo.CollectProductCount,
		CollectSubjectCount: c.UmsMemberStatisticsInfo.CollectSubjectCount,
		CollectTopicCount:   c.UmsMemberStatisticsInfo.CollectTopicCount,
		CollectCommentCount: c.UmsMemberStatisticsInfo.CollectCommentCount,
		InviteFriendCount:   c.UmsMemberStatisticsInfo.InviteFriendCount,
		RecentOrderTime:     c.UmsMemberStatisticsInfo.ReturnOrderCount,
	}
	umsMemberStatisticsInfoId, err := o.Insert(&umsMemberStatisticsInfo)
	if err != nil {
		o.Begin()
	}

	// 插入字表
	// 基本实例化 UmsMember 结构体
	umsMember := UmsMember{
		UserName:                c.UserName,
		Password:                c.Password,
		NickName:                c.NickName,
		Phone:                   c.Phone,
		Status:                  c.Status,
		CreateTime:              c.CreateTime,
		Icon:                    c.Icon,
		Gender:                  c.Gender,
		Birthday:                c.Birthday,
		City:                    c.City,
		Job:                     c.Job,
		PersonalizedSignature:   c.PersonalizedSignature,
		SourceType:              c.SourceType,
		Integration:             c.Integration,
		Growth:                  c.Growth,
		LuckeyCount:             c.LuckeyCount,
		HistoryIntegration:      c.HistoryIntegration,
		UmsMemberLevel:          &UmsMemberLevel{Id: int(umsMemberLevelId)},
		UmsMemberStatisticsInfo: &UmsMemberStatisticsInfo{Id: int(umsMemberStatisticsInfoId)},
	}

	id, err = o.Insert(&umsMember)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}
	// 提交事务
	err = o.Commit()
	return id, err
}

// AddUmsMember 新增
func AddUmsMember(c UmsMember) (id int64, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 开启事务
	o.Begin()
	// 插入主表 UmsMemberLevel
	umsMemberLevel := UmsMemberLevel{
		Name:                  c.UmsMemberLevel.Name,
		GrowthPoint:           c.UmsMemberLevel.GrowthPoint,
		DefaultStatus:         c.UmsMemberLevel.DefaultStatus,
		FreeFreightPoint:      c.UmsMemberLevel.FreeFreightPoint,
		CommentGrowthPoint:    c.UmsMemberLevel.CommentGrowthPoint,
		PriviledgeFreeFreight: c.UmsMemberLevel.PriviledgeFreeFreight,
		PriviledgeSifnIn:      c.UmsMemberLevel.PriviledgeSifnIn,
		PriviledgeComment:     c.UmsMemberLevel.PriviledgeComment,
		PriviledgePromotion:   c.UmsMemberLevel.PriviledgePromotion,
		PriviledgeMemberPrice: c.UmsMemberLevel.PriviledgeMemberPrice,
		PriviledgeBirthday:    c.UmsMemberLevel.PriviledgeBirthday,
		Note:                  c.UmsMemberLevel.Note,
	}
	umsMemberLevelId, err := o.Insert(&umsMemberLevel)
	if err != nil {
		// 回滚事务
		err = o.Begin()
	}

	// 插入主表 UmsMemberStatisticsInfo
	umsMemberStatisticsInfo := UmsMemberStatisticsInfo{
		ConsumeAmount:       c.UmsMemberStatisticsInfo.ConsumeAmount,
		OrderCount:          c.UmsMemberStatisticsInfo.OrderCount,
		CouponCount:         c.UmsMemberStatisticsInfo.CouponCount,
		CommentCount:        c.UmsMemberStatisticsInfo.CommentCount,
		ReturnOrderCount:    c.UmsMemberStatisticsInfo.ReturnOrderCount,
		LoginCount:          c.UmsMemberStatisticsInfo.LoginCount,
		AttendCount:         c.UmsMemberStatisticsInfo.AttendCount,
		FansCount:           c.UmsMemberStatisticsInfo.FansCount,
		CollectProductCount: c.UmsMemberStatisticsInfo.CollectProductCount,
		CollectSubjectCount: c.UmsMemberStatisticsInfo.CollectSubjectCount,
		CollectTopicCount:   c.UmsMemberStatisticsInfo.CollectTopicCount,
		CollectCommentCount: c.UmsMemberStatisticsInfo.CollectCommentCount,
		InviteFriendCount:   c.UmsMemberStatisticsInfo.InviteFriendCount,
		RecentOrderTime:     c.UmsMemberStatisticsInfo.ReturnOrderCount,
	}
	umsMemberStatisticsInfoId, err := o.Insert(&umsMemberStatisticsInfo)
	if err != nil {
		o.Begin()
	}

	// 插入字表
	// 基本实例化 UmsMember 结构体
	umsMember := UmsMember{
		UserName:                c.UserName,
		Password:                c.Password,
		NickName:                c.NickName,
		Phone:                   c.Phone,
		Status:                  c.Status,
		CreateTime:              c.CreateTime,
		Icon:                    c.Icon,
		Gender:                  c.Gender,
		Birthday:                c.Birthday,
		City:                    c.City,
		Job:                     c.Job,
		PersonalizedSignature:   c.PersonalizedSignature,
		SourceType:              c.SourceType,
		Integration:             c.Integration,
		Growth:                  c.Growth,
		LuckeyCount:             c.LuckeyCount,
		HistoryIntegration:      c.HistoryIntegration,
		UmsMemberLevel:          &UmsMemberLevel{Id: int(umsMemberLevelId)},
		UmsMemberStatisticsInfo: &UmsMemberStatisticsInfo{Id: int(umsMemberStatisticsInfoId)},
	}

	id, err = o.Insert(&umsMember)
	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}
	// 提交事务
	err = o.Commit()
	return id, err
}

// UpdateUmsMember 更新
func UpdateUmsMember(uid int, uu *UmsMember) (a *UmsMember, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 创建 UmsMemberLevel 对象
	umsMemberLevel := UmsMemberLevel{Id: uu.UmsMemberLevel.Id}
	// 创建 UmsMemberStatisticsInfo 对象
	umsMemberStatisticsInfo := UmsMemberStatisticsInfo{Id: uu.UmsMemberStatisticsInfo.Id}

	umsMemberLevelErr := o.Read(&umsMemberLevel)
	umsMemberStatisticsInfoErr := o.Read(&umsMemberStatisticsInfo)
	if umsMemberLevelErr == nil && umsMemberStatisticsInfoErr == nil {
		umsMember := UmsMember{
			Id:                      uid,
			UmsMemberLevel:          &umsMemberLevel,
			UmsMemberStatisticsInfo: &umsMemberStatisticsInfo,
			UserName:                uu.UserName,
			Password:                uu.Password,
			NickName:                uu.NickName,
			Phone:                   uu.Phone,
			Status:                  uu.Status,
			CreateTime:              uu.CreateTime,
			Icon:                    uu.Icon,
			Gender:                  uu.Gender,
			Birthday:                uu.Birthday,
			City:                    uu.City,
			Job:                     uu.Job,
			PersonalizedSignature:   uu.PersonalizedSignature,
			SourceType:              uu.SourceType,
			Integration:             uu.Integration,
			Growth:                  uu.Growth,
			LuckeyCount:             uu.LuckeyCount,
			HistoryIntegration:      uu.HistoryIntegration,
		}
		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&umsMember); err != nil {
			return nil, errors.New("修改失败")
		}

		if err != nil {
			// 事务回退
			err = o.Rollback()
		} else {
			// 提交事务
			err = o.Commit()
		}
		return &umsMember, nil
	}
	return nil, err
}

// LoginOfUmsMember 会员登录
func LoginOfUmsMember(u UmsMember) (c *UmsMember, err error) {
	o := orm.NewOrm()
	umsMember := UmsMember{Phone: u.Phone, Password: u.Password}

	err = o.Read(&umsMember, "Phone", "Password")

	if err == nil {
		if umsMember.UmsMemberLevel != nil {
			err = o.Read(umsMember.UmsMemberLevel)
		}
		if umsMember.UmsMemberStatisticsInfo != nil {
			err = o.Read(umsMember.UmsMemberStatisticsInfo)
		}
		return &umsMember, err
	}
	return &umsMember, err
}

// LogoutOfUmsMember 退出登录
func LogoutOfUmsMember(u UmsMember) (c *UmsMember, err error) {
	o := orm.NewOrm()
	umsMember := UmsMember{Phone: u.Phone, Password: u.Password}

	err = o.Read(&umsMember, "Phone", "Password")

	if err == nil {
		return &umsMember, err
	}
	return &umsMember, err
}

// GetUmsMember 查询单个
func GetUmsMember(uid int) (c *UmsMember, err error) {
	o := orm.NewOrm()
	umsMember := UmsMember{Id: uid}

	err = o.Read(&umsMember)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		if umsMember.UmsMemberLevel != nil {
			err = o.Read(umsMember.UmsMemberLevel)
		}
		if umsMember.UmsMemberStatisticsInfo != nil {
			err = o.Read(umsMember.UmsMemberStatisticsInfo)
		}

	}

	return &umsMember, err

}

// GetAllUmsMember 分页查询评论
func GetAllUmsMember(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()

	var pmsProduct []*UmsMember

	qs := o.QueryTable("ums_member")

	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&pmsProduct)
	for _, u := range pmsProduct {
		if u.UmsMemberLevel != nil {
			err = o.Read(u.UmsMemberLevel)
		}
		if u.UmsMemberStatisticsInfo != nil {
			err = o.Read(u.UmsMemberStatisticsInfo)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, pmsProduct), err
}

// DeleteUmsMember 删除指定评论
func DeleteUmsMember(uid int) (b bool, err error) {
	// 创建 oremer 实例
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除表
	umsMember := UmsMember{Id: uid}
	_, err = o.Delete(&umsMember)

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
	orm.RegisterModel(new(UmsMember))
}
