package main

import (
	"chat-application/initializers"
	_ "chat-application/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {
	//sessionConf := &session.ManagerConfig{
	//	CookieName: "beegoSessionID",
	//	Gclifetime: 3600,
	//}
	//beego.GlobalSessions,_ = session.NewManager("memory", sessionConf)
	//go beego.GlobalSessions.GC()
	beego.Run()
}
