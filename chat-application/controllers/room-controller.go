package controllers

import (
	"chat-application/initializers"
	users "chat-application/sqlc-models"
	"chat-application/utils"
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type RoomController struct {
	beego.Controller
}

// URLMapping ...
func (mapping *RoomController) URLMapping() {
	mapping.Mapping("CreateRoom", mapping.CreateRoom)
	mapping.Mapping("GetAllRooms", mapping.GetAllRooms)
}

// CreateRoom  ...
// @Title Post
// @Description create room
// @Param	body		body 	users.Room	true		"body for add room content"
// @Success 201 {object} users.Room
// @Failure 403 body is empty
// @router /create [post]
func (room *RoomController) CreateRoom() {
	var (
		roomObj users.Room
	)
	if room.Ctx.Input.Method() == "POST" {
		roomObj = room.FormParser(roomObj)
		result, createError := initializers.Db.CreateRoom(context.Background(), users.CreateRoomParams{
			RoomName: roomObj.RoomName,
			UserID:   credentials.ID,
		})
		logs.Info(result)
		utils.CheckError(createError, "Error creating records")

	}
}

// GetAllRooms  ...
// @Title Get
// @Description Get room
// @Success 200 {object} users.Room
// @Failure 403
// @router / [get]
func (room *RoomController) GetAllRooms() {
	var (
		roomObj  []users.GetRoomsRow
		getError error
		found    bool
	)
	roomObj, getError = initializers.Db.GetRooms(context.Background())
	if getError != nil {
		logs.Info("Error getting records: Reason ", getError)
		return
	}
	if len(roomObj) == 0 {
		found = false
	} else {
		found = true
	}
	room.Data["Found"] = found
	room.Data["Rooms"] = roomObj
	logs.Info(credentials.ID)
	room.Data["FullName"] = credentials.FullName
	room.TplName = "default/home-page.html"
}
func (room *RoomController) FormParser(roomObj users.Room) users.Room {
	//parseErr := room.ParseForm(&roomObj)
	//utils.CheckError(parseErr, "Error parsing form")
	if room.Ctx.Input.Method() == "POST" {
		roomObj.RoomName = room.GetString("room")
		return roomObj
	}
	return users.Room{}
}
