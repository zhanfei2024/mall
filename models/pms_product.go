package models

import (
	"errors"
	"mall/utils"

	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// PmsProduct 定义结构体
type PmsProduct struct {
	Id                          int                          `json:"id"`
	Name                        string                       `description:"商品名称" orm:"size(64)" json:"name"`
	Pic                         string                       `description:"图片" json:"pic"`
	ProductSn                   string                       `description:"货号" orm:"size(64)" json:"product_sn"`
	DeleteStatus                int                          `description:"删除状态：0->未删除；1->已删除" json:"delete_status"`
	PublishStatus               int                          `description:"上架状态：0->下架；1->上架" json:"publish_status"`
	NewStatus                   int                          `description:"新品状态:0->不是新品；1->新品" json:"new_status"`
	RecommendStatus             int                          `description:"推荐状态；0->不推荐；1->推荐" json:"recommend_status"`
	VerifyStatus                int                          `description:"审核状态：0->未审核；1->审核通过" json:"verify_status"`
	Sort                        int                          `description:"排序" json:"sort"`
	Sale                        int                          `description:"销量" json:"sale"`
	Price                       float64                      `orm:"digits(10);decimals(2)" description:"价格" json:"price"`
	PromotionPrice              float64                      `orm:"digits(10);decimals(2)" description:"促销价格" json:"price"`
	GiftGrowth                  int                          `orm:"default(0)" description:"赠送的成长值" json:"gift_growth"`
	GiftPoint                   int                          `orm:"default(0)" description:"赠送的积分" json:"gift_point"`
	UsePointLimit               int                          `description:"限制使用的积分数" json:"use_point_limit"`
	SubTitle                    string                       `description:"副标题" json:"sub_title"`
	Description                 string                       `description:"商品描述" json:"description"`
	OriginalPrice               float64                      `orm:"digits(10);decimals(2)" description:"市场价" json:"original_price"`
	Stock                       int                          `description: "库存" json:"stock"`
	LowStock                    int                          `description: "库存预警值" json:"low_stock"`
	Unit                        string                       `orm:size(16) description:"单位" json:"unit"`
	Weight                      float64                      `orm:"digits(10);decimals(2)" description:"商品重量" json:"weight"`
	PreviewStatus               int                          `description:"是否为预告商品" json:"preview_status"`
	ServiceIds                  string                       `orm:size(64) description:"以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮" json:"service_ids"`
	KeyWords                    string                       `description:"关键字" json:"KeyWords"`
	Note                        string                       `description:"备注" json:"note"`
	AlbumPics                   string                       `description:"画册图片，连产品图片限制为5张，以逗号分割" json:"album_pics"`
	DetailTitle                 string                       `description:"详情标题" json:"detail_title"`
	DetailDesc                  string                       `description:"详情描述" json:"detail_desc"`
	DetailHtml                  string                       `description:"产品详情网页内容" json:"detail_html"`
	DetailMobileHtml            string                       `description:"移动端网页详情" json:"detail_mobile_html"`
	PromotionStarTime           time.Time                    `orm:"auto_now_add;type(datetime)" description:"促销开始时间" json:"promotion_start_time"`
	PromotionEndTime            time.Time                    `orm:"auto_now_add;type(datetime)" description:"促销结束时间" json:"promotion_end_time"`
	PromotionPerLimit           int                          `description:"活动限购数量" json:"promotion_per_limit"`
	PromotionType               int                          `description:"促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购" json:"promotion_type"`
	ProductGategoryName         string                       `description:"产品分类名称" json:"promotion_type"`
	BrandName                   string                       `description:"品牌名称" json:"brand_name"`
	PmsBrand                    *PmsBrand                    `orm:"rel(fk)" json:"pms_brand"`
	PmsProductCategory          *PmsProductCategory          `orm:"rel(fk)" json:"pms_product_category"`
	PmsProductAttributeCategory *PmsProductAttributeCategory `orm:"rel(fk)" json:"pms_product_attribute_category"`
	PmsFeightTemplate           *PmsFeightTemplate           `orm:"rel(fk)" json:"pms_feight_template"`
}

// TableName 自定义表名
func (c *PmsProduct) TableName() string {
	return "pms_product"
}

// AddPmsProduct 新增
func AddPmsProduct(c PmsProduct) (id int64, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 基本实例化 pmsBrand 结构体
	pmsBrand := PmsBrand{Id: c.PmsBrand.Id}
	// 基本实例化 pmsProductCategory 结构体
	pmsProductCategory := PmsProductCategory{Id: c.PmsProductCategory.Id}
	// 基本实例化 pmsProductAttributeCategory 结构体
	pmsProductAttributeCategory := PmsProductAttributeCategory{Id: c.PmsProductAttributeCategory.Id}
	// 基本实例化 PmsFeightTemplate 结构体
	pmsFeightTemplate := PmsFeightTemplate{Id: c.PmsFeightTemplate.Id}
	// 基本实例化 pmsProduct 结构体
	pmsProduct := PmsProduct{
		Name:                c.Name,
		Pic:                 c.Pic,
		ProductSn:           c.ProductSn,
		DeleteStatus:        c.DeleteStatus,
		PublishStatus:       c.PublishStatus,
		NewStatus:           c.NewStatus,
		RecommendStatus:     c.RecommendStatus,
		VerifyStatus:        c.VerifyStatus,
		Sort:                c.Sort,
		Sale:                c.Sale,
		Price:               c.Price,
		PromotionPrice:      c.PromotionPrice,
		GiftGrowth:          c.GiftGrowth,
		GiftPoint:           c.GiftPoint,
		UsePointLimit:       c.UsePointLimit,
		SubTitle:            c.SubTitle,
		Description:         c.Description,
		OriginalPrice:       c.OriginalPrice,
		Stock:               c.Stock,
		LowStock:            c.LowStock,
		Unit:                c.Unit,
		Weight:              c.Weight,
		PreviewStatus:       c.PreviewStatus,
		ServiceIds:          c.ServiceIds,
		KeyWords:            c.KeyWords,
		Note:                c.Note,
		AlbumPics:           c.AlbumPics,
		DetailTitle:         c.DetailTitle,
		DetailHtml:          c.DetailHtml,
		DetailMobileHtml:    c.DetailMobileHtml,
		PromotionStarTime:   time.Now(),
		PromotionEndTime:    time.Now(),
		PromotionPerLimit:   c.PromotionPerLimit,
		PromotionType:       c.PromotionType,
		ProductGategoryName: c.ProductGategoryName,
		BrandName:           c.BrandName,
	}

	// &pmsBrand  指针类型的结构体
	psmBrandErr := o.Read(&pmsBrand)

	// &pmsProductCategory 指针类型的结构体
	pmsProductCategoryErr := o.Read(&pmsProductCategory)

	// &pmsProductAttributeCategory 指针类型的结构体
	pmsProductAttributeCategoryErr := o.Read(&pmsProductAttributeCategory)

	// &pmsFeightTemplate 指针类型的结构体
	pmsFeightTemplateErr := o.Read(&pmsFeightTemplate)

	if psmBrandErr == orm.ErrNoRows || pmsProductCategoryErr == orm.ErrNoRows || pmsProductAttributeCategoryErr == orm.ErrNoRows || pmsFeightTemplateErr == orm.ErrNoRows {
		fmt.Println("该商品的品牌或分类或属性或运费模板不存在")
	} else if psmBrandErr == orm.ErrMissPK || pmsProductCategoryErr == orm.ErrMissPK || pmsProductAttributeCategoryErr == orm.ErrMissPK || pmsFeightTemplateErr == orm.ErrMissPK {
		fmt.Println("该商品的品牌或分类或属性或运费模板的主键不存在")
	} else {
		pmsProduct.PmsBrand = &pmsBrand
		pmsProduct.PmsProductCategory = &pmsProductCategory
		pmsProduct.PmsProductAttributeCategory = &pmsProductAttributeCategory
		pmsProduct.PmsFeightTemplate = &pmsFeightTemplate

		// 开启事务
		o.Begin()

		// 插入评论
		id, err = o.Insert(&pmsProduct)
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

// UpdatePmsProduct 更新
func UpdatePmsProduct(uid int, uu *PmsProduct) (a *PmsProduct, err error) {
	// 创建 ormer 实例
	o := orm.NewOrm()
	// 创建 PmsBrand 对象
	pmsBrand := PmsBrand{Id: uu.PmsBrand.Id}
	// 创建 pmsProductCategory 对象
	pmsProductCategory := PmsProductCategory{Id: uu.PmsProductCategory.Id}
	// 创建 pmsProductAttributeCategory 对象
	pmsProductAttributeCategory := PmsProductAttributeCategory{Id: uu.PmsProductAttributeCategory.Id}
	// 创建 PmsFeightTemplate 对象
	pmsFeightTemplate := PmsFeightTemplate{Id: uu.PmsFeightTemplate.Id}

	psmBrandErr := o.Read(&pmsBrand)

	pmsProductCategoryErr := o.Read(&pmsProductCategory)

	pmsProductAttributeCategoryErr := o.Read(&pmsProductAttributeCategory)

	pmsFeightTemplateErr := o.Read(&pmsFeightTemplate)

	if psmBrandErr == nil && pmsProductCategoryErr == nil && pmsProductAttributeCategoryErr == nil && pmsFeightTemplateErr == nil {
		pmsProduct := PmsProduct{
			Id:                          uid,
			PmsBrand:                    &pmsBrand,
			PmsProductCategory:          &pmsProductCategory,
			PmsProductAttributeCategory: &pmsProductAttributeCategory,
			PmsFeightTemplate:           &pmsFeightTemplate,
			Name:                        uu.Name,
			Pic:                         uu.Pic,
			ProductSn:                   uu.ProductSn,
			DeleteStatus:                uu.DeleteStatus,
			PublishStatus:               uu.PublishStatus,
			NewStatus:                   uu.NewStatus,
			RecommendStatus:             uu.RecommendStatus,
			VerifyStatus:                uu.VerifyStatus,
			Sort:                        uu.Sort,
			Sale:                        uu.Sale,
			Price:                       uu.Price,
			PromotionPrice:              uu.PromotionPrice,
			GiftGrowth:                  uu.GiftGrowth,
			GiftPoint:                   uu.GiftPoint,
			UsePointLimit:               uu.UsePointLimit,
			SubTitle:                    uu.SubTitle,
			Description:                 uu.Description,
			OriginalPrice:               uu.OriginalPrice,
			Stock:                       uu.Stock,
			LowStock:                    uu.LowStock,
			Unit:                        uu.Unit,
			Weight:                      uu.Weight,
			PreviewStatus:               uu.PreviewStatus,
			ServiceIds:                  uu.ServiceIds,
			KeyWords:                    uu.KeyWords,
			Note:                        uu.Note,
			AlbumPics:                   uu.AlbumPics,
			DetailTitle:                 uu.DetailTitle,
			DetailHtml:                  uu.DetailHtml,
			DetailMobileHtml:            uu.DetailMobileHtml,
			PromotionStarTime:           time.Now(),
			PromotionEndTime:            time.Now(),
			PromotionPerLimit:           uu.PromotionPerLimit,
			PromotionType:               uu.PromotionType,
			ProductGategoryName:         uu.ProductGategoryName,
			BrandName:                   uu.BrandName,
		}
		// 开启事务
		err = o.Begin()

		if _, err := o.Update(&pmsProduct); err != nil {
			return nil, errors.New("修改失败")
		}

		if err != nil {
			// 事务回退
			err = o.Rollback()
		} else {
			// 提交事务
			err = o.Commit()
		}
		return &pmsProduct, nil
	}
	return nil, err
}

// GetPmsProduct 查询单个
func GetPmsProduct(uid int) (c *PmsProduct, err error) {
	o := orm.NewOrm()
	pmsProduct := PmsProduct{Id: uid}

	err = o.Read(&pmsProduct)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		if pmsProduct.PmsBrand != nil {
			err = o.Read(pmsProduct.PmsBrand)
		}
		if pmsProduct.PmsProductCategory != nil {
			err = o.Read(pmsProduct.PmsProductCategory)
		}
		if pmsProduct.PmsProductAttributeCategory != nil {
			err = o.Read(pmsProduct.PmsProductAttributeCategory)
		}
		if pmsProduct.PmsFeightTemplate != nil {
			err = o.Read(pmsProduct.PmsFeightTemplate)
		}

	}

	return &pmsProduct, err

}

// GetAllPmsProducts 分页查询评论
func GetAllPmsProducts(p int, size int) (u utils.Page, err error) {
	o := orm.NewOrm()

	var pmsProduct []*PmsProduct

	qs := o.QueryTable("pms_product")

	count, _ := qs.Limit(-1).Count()
	_, err = qs.RelatedSel().Limit(size).Offset((p - 1) * size).All(&pmsProduct)
	for _, u := range pmsProduct {
		if u.PmsBrand != nil {
			err = o.Read(u.PmsBrand)
		}
		if u.PmsProductCategory != nil {
			err = o.Read(u.PmsProductCategory)
		}
		if u.PmsProductAttributeCategory != nil {
			err = o.Read(u.PmsProductAttributeCategory)
		}
		if u.PmsFeightTemplate != nil {
			err = o.Read(u.PmsFeightTemplate)
		}
	}
	c, _ := strconv.Atoi(strconv.FormatInt(count, 10))
	return utils.Pagination(c, p, size, pmsProduct), err
}

// DeletePmsProduct 删除指定评论
func DeletePmsProduct(uid int) (b bool, err error) {
	// 创建 oremer 实例
	o := orm.NewOrm()

	// 开启事务
	err = o.Begin()

	// 删除表
	pmsProduct := PmsProduct{Id: uid}
	_, err = o.Delete(&pmsProduct)

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
	orm.RegisterModel(new(PmsProduct))
}
