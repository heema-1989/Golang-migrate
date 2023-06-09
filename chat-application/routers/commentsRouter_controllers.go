package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

	beego.GlobalControllerRouter["chat-application/controllers:ChatController"] = append(beego.GlobalControllerRouter["chat-application/controllers:ChatController"],
		beego.ControllerComments{
			Method:           "GetAllChat",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:LoginController"] = append(beego.GlobalControllerRouter["chat-application/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:LoginController"] = append(beego.GlobalControllerRouter["chat-application/controllers:LoginController"],
		beego.ControllerComments{
			Method:           "LoggedIn",
			Router:           "/home",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:RegisterController"] = append(beego.GlobalControllerRouter["chat-application/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "RegisterUser",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:RegisterController"] = append(beego.GlobalControllerRouter["chat-application/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "VerifyUser",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:RegisterController"] = append(beego.GlobalControllerRouter["chat-application/controllers:RegisterController"],
		beego.ControllerComments{
			Method:           "VerifiedUser",
			Router:           "/verified",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:RoomController"] = append(beego.GlobalControllerRouter["chat-application/controllers:RoomController"],
		beego.ControllerComments{
			Method:           "GetAllRooms",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["chat-application/controllers:RoomController"] = append(beego.GlobalControllerRouter["chat-application/controllers:RoomController"],
		beego.ControllerComments{
			Method:           "CreateRoom",
			Router:           "/create",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
