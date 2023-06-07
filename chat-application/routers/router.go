// @APIVersion 1.0.0
// @Title Case Study Builder
// @Description List Of All Case Study API's
// @Contact -
// @TermsOfServiceUrl -
// @License Apache -
// @LicenseUrl -
package routers

import (
	"chat-application/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/register", &controllers.RegisterController{}, "get:RegisterUser")
	//beego.Router("/verify-register", &controllers.RegisterController{}, "post:VerifyUser")
	//beego.Router("/verified", &controllers.RegisterController{}, "get:VerifiedUser")
	//beego.Router("/login", &controllers.LoginController{}, "*:Login")
	//beego.Router("/home", &controllers.LoginController{}, "post:LoggedIn")
	//beego.Router("/room/create", &controllers.RoomController{}, "post:CreateRoom")
	//beego.Router("/room/get", &controllers.RoomController{}, "get:GetAllRooms")
	nsApi := beego.NewNamespace("/api",
		beego.NSNamespace("/register", beego.NSInclude(&controllers.RegisterController{})),
		beego.NSNamespace("/login", beego.NSInclude(&controllers.LoginController{})),
		beego.NSNamespace("/home", beego.NSInclude(&controllers.LoginController{})),
	)
	beego.AddNamespace(nsApi)
}
