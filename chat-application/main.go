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
	beego.Run()
}
