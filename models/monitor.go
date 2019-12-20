package models

import (
	"errors"
	"first_go/utils"
	"strconv"

	"github.com/astaxie/beego/orm"
)

// Monitor 监视器结构体
type Monitor struct {
	Id     int     `json:"id"`
	Lon    float32 `json:"lon"`
	Lat    float32 `json:"lat"`
	Addr   string  `json:"addr"`
	Status int     `json:"status"`
}

// TableName 自定义表名
func (u *Monitor) TableName() string {
	return "monitor"
}

// AddMonitor 新增
func AddMonitor(u Monitor) (id int64, err error) {
	// one2one 插入
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 插入
	monitor := Monitor{
		Id:     u.Id,
		Lon:    u.Lon,
		Lat:    u.Lat,
		Addr:   u.Addr,
		Status: u.Status,
	}

	id, err = o.Insert(&monitor)

	if err != nil {
		// 回滚事务
		err = o.Rollback()
	}

	// 提交事务
	err = o.Commit()
	return id, err
}

// GetMonitor 查询
func GetMonitor(uid int) (m *Monitor, err error) {
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	monitor := &Monitor{Id: uid}
	err = o.Read(monitor)

	return monitor, err
}

// GetAllMonitor 分页查询
func GetAllMonitor(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()
	monitor := new(Monitor)
	var monitors []Monitor
	qs := o.QueryTable(monitor)
	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&monitors)
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, monitors), err
}

// UpdateMonitor 更新指定
func UpdateMonitor(uid int, mm *Monitor) (a *Monitor, err error) {
	o := orm.NewOrm()
	monitor := Monitor{Id: uid}

	if o.Read(&monitor) == nil {

		if mm.Lon != 0 {
			monitor.Lon = mm.Lon
		}
		if mm.Lat != 0 {
			monitor.Lat = mm.Lat
		}
		if mm.Addr != "" {
			monitor.Addr = mm.Addr
		}
		if mm.Status != 0 {
			monitor.Status = mm.Status
		}

		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&monitor); err != nil {
			return nil, errors.New("修改失败")
		}

		if err != nil {
			err = o.Rollback()
		} else {
			err = o.Commit()
		}
		return &monitor, nil
	}

	return nil, err
}

// DeleteMonitor 删除指定用户
func DeleteMonitor(uid int) (b bool, err error) {
	// one2one 删除
	// 创建一个 ormer 对象
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除
	monitor := Monitor{Id: uid}
	_, err = o.Delete(&monitor)
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
	orm.RegisterModel(new(Monitor))

}
