package controllers

import (
	"chat-application/initializers"
	"chat-application/models"
	users "chat-application/sqlc-models"
	"chat-application/utils"
	"context"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type RoomController struct {
	beego.Controller
}

func (room *RoomController) CreateRoom() {
	var (
		roomObj users.Room
		resp    models.JsonResponse
	)
	if room.Ctx.Input.Method() == "POST" {
		unmarshalErr := json.Unmarshal(room.Ctx.Input.RequestBody, &roomObj)
		if unmarshalErr != nil {
			resp = models.Response(406, "Something went wrong parsing data", make([]string, 0), resp)
		} else {
			result, createError := initializers.Db.CreateRoom(context.Background(), users.CreateRoomParams{
				RoomName: roomObj.RoomName,
				UserID:   roomObj.UserID,
			})
			utils.CheckError(createError, "Error creating records")
			resp = models.Response(200, "Room successfully created", result, resp)
		}
		Send(&room.Controller, resp)
	}
}

func (room *RoomController) GetAllRooms() {
	var (
		roomObj  []users.GetRoomsRow
		resp     models.JsonResponse
		getError error
		found    bool
	)
	roomObj, getError = initializers.Db.GetRooms(context.Background())
	if getError != nil {
		logs.Info("Error getting records: Reason ", getError)
		return
	}
	if len(roomObj) == 0 {
		resp = models.Response(200, "Records not found", make([]string, 0), resp)
	} else {
		found = true
		resp = models.Response(200, "Records have been successfully retrieved", roomObj, resp)
	}
	room.Data["Found"] = found
	room.Data["Rooms"] = resp.Data
	//err := utils.ValidateJwt(credentials, credentials.JwtToken, loginKey)
	//if err != nil {
	//	logs.Info("Token expired")
	//	return
	//}
	//room.Data["FullName"] = credentials.FullName
	room.TplName = "default/home-page.html"
	//Send(&room.Controller, resp)
}
