package controllers

import (
	"chat-application/models"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func Send(controller *beego.Controller, resp models.JsonResponse) {
	controller.Data["json"] = resp
	serveJsonError := controller.ServeJSON()
	if serveJsonError != nil {
		logs.Error("Error serving json: Reason ", serveJsonError)
		controller.Abort("500")
	}
	controller.Finish()
}
