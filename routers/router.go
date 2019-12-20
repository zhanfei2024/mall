package routers

// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html

import (
	"mall/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/register",
			beego.NSInclude(
				&controllers.RegisterController{},
			),
		),
		beego.NSNamespace("/login",
			beego.NSInclude(
				&controllers.LoginController{},
			),
		),
		beego.NSNamespace("/logout",
			beego.NSInclude(
				&controllers.LogoutController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/monitor",
			beego.NSInclude(
				&controllers.MonitorController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.CommentController{},
			),
		),
		beego.NSNamespace("/post",
			beego.NSInclude(
				&controllers.PostController{},
			),
		),
		beego.NSNamespace("/tag",
			beego.NSInclude(
				&controllers.TagController{},
			),
		),
		beego.NSNamespace("/pms_product",
			beego.NSInclude(
				&controllers.PmsProductController{},
			),
		),
		beego.NSNamespace("/pms_product_category",
			beego.NSInclude(
				&controllers.PmsProductCategoryController{},
			),
		),
		beego.NSNamespace("/pms_brand",
			beego.NSInclude(
				&controllers.PmsBrandController{},
			),
		),
		beego.NSNamespace("/pms_product_attribute_category",
			beego.NSInclude(
				&controllers.PmsProductAttributeCategoryController{},
			),
		),
		beego.NSNamespace("/pms_feight_template",
			beego.NSInclude(
				&controllers.PmsFeightTemplateController{},
			),
		),
		beego.NSNamespace("/ums_member",
			beego.NSInclude(
				&controllers.UmsMemberController{},
			),
		),
		beego.NSNamespace("/ums_member_level",
			beego.NSInclude(
				&controllers.UmsMemberLevelController{},
			),
		),
		beego.NSNamespace("/ums_member_statistics_info",
			beego.NSInclude(
				&controllers.UmsMemberStatisticsInfoController{},
			),
		),
		beego.NSNamespace("/ums_member_receive_address",
			beego.NSInclude(
				&controllers.UmsMemberReceiveAddressController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
