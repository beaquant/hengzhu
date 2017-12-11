package routers

import (
	"github.com/astaxie/beego"
	"hengzhu/controllers"
)

func init() {
	// 默认登录
	beego.Router("/", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login", &controllers.LoginController{}, "*:LoginIn")
	beego.Router("/login_out", &controllers.LoginController{}, "*:LoginOut")
	beego.Router("/no_auth", &controllers.LoginController{}, "*:NoAuth")

	beego.Router("/home", &controllers.HomeController{}, "*:Index")
	beego.Router("/home/start", &controllers.HomeController{}, "*:Start")

	beego.Router("/cabinetDetail/table", &controllers.CabinetDetailController{}, "*:Table")
	beego.Router("/cabinetDetail/edit", &controllers.CabinetDetailController{}, "*:Table")
	//预下单
	beego.Router("/order/reorder", &controllers.OrderController{}, "*:ReOrder")
	//支付通知
	beego.Router("/pay_notify/ali", &controllers.PayNotifyController{},"*:AliNotify")
	beego.Router("/pay_notify/wx", &controllers.PayNotifyController{},"*:WxNotify")
	//beego.Router("/setting/get", &controllers.SettingController{}, "*:Get")

	beego.AutoRouter(&controllers.CabinetController{})
	beego.AutoRouter(&controllers.CabinetDetailController{})
	beego.AutoRouter(&controllers.TypesController{})
	beego.AutoRouter(&controllers.SettingController{})
	beego.AutoRouter(&controllers.AuthController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.AdminController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.PayNotifyController{})

	//beego.Router("/account", &controllers.AccountController{}, "*:Index")
	//beego.Router("/account/changepwd", &controllers.AccountController{}, "*:ChangePwd")
	//beego.Router("/account/bind", &controllers.AccountController{}, "*:Bind")
	//beego.Router("/account/updatedetail", &controllers.AccountController{}, "*:UpdateDetail")
	//beego.Router("/sendvcode", &controllers.UserController{}, "post:SendVCode")
	//beego.Router("/service", &controllers.ServiceController{}, "*:Service")
	//beego.Router("/detail", &controllers.ServiceController{}, "*:Detail")
	//admin.Run()
	//ucenter.Run()
}
