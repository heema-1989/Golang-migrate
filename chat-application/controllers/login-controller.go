package controllers

import (
	"chat-application/initializers"
	users "chat-application/sqlc-models"
	"chat-application/utils"
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"strings"
)

type LoginController struct {
	beego.Controller
}

var (
	credentials users.User
	flag, check bool = true, true
	loginKey    []byte
)

// URLMapping ...
func (mapping *LoginController) URLMapping() {
	mapping.Mapping("LoginUser", mapping.Login)
	mapping.Mapping("HomePage", mapping.LoggedIn)
}

// Login  ...
// @Title Get
// @Description get Login Page
// @Success 200 {object} models.Token
// @Failure 403
// @router / [get]
func (login *LoginController) Login() {
	login.Data["Flag"] = flag
	login.Data["Check"] = check
	if !flag {
		if !check {
			login.SetCredentialsTemplateNames(credentials)
		}
	}
	login.TplName = "default/login.html"
	flag = true
	check = true
	println("djkddd")
}

// LoggedIn  ...
// @Title Get
// @Description get Home Page
// @Success 200 {object} models.Token
// @Failure 403
// @router /home [post]
func (login *LoginController) LoggedIn() {

	credentials = login.FormParser(credentials)
	result, getErr := initializers.Db.GetUserEmail(context.Background(), credentials.Email)
	if !login.CheckExists(getErr) {
		flag = false
		login.Data["Flag"] = flag
		login.Data["Check"] = check
		login.Redirect("http://localhost:8080/api/login", 302)
		return
	}
	creds, passErr := initializers.Db.GetUserCredentials(context.Background(), users.GetUserCredentialsParams{
		Email:    result.Email,
		Password: credentials.Password,
	})
	if !login.CheckExists(passErr) {
		flag = false
		check = false
		login.Data["Flag"] = flag
		login.Data["Check"] = check
		login.Redirect("http://localhost:8080/api/login", 302)
		return
	}
	login.SetCredentialsTemplateNames(creds)
	logs.Info(creds)
	credentials = creds
	login.Redirect("http://localhost:8080/api/home", 302)
}
func (login *LoginController) SetCredentialsTemplateNames(credentials users.User) {
	login.Data["VerifyId"] = credentials.VerifyID
	login.Data["UserName"] = credentials.UserName
	login.Data["FullName"] = credentials.FullName
	login.Data["Email"] = credentials.Email
	login.Data["Password"] = credentials.Password
}

func (login *LoginController) FormParser(credentials users.User) users.User {
	parseErr := login.ParseForm(&credentials)
	utils.CheckError(parseErr, "Error parsing form")
	if login.Ctx.Input.Method() == "POST" {
		credentials.Email = login.GetString("email")
		credentials.Password = login.GetString("password")
		return credentials
	}
	return users.User{}
}
func (login *LoginController) CheckExists(err error) bool {
	if err != nil {
		if strings.Contains(err.Error(), "sql: no rows in result set") {

			return false
		}
		logs.Info("Error: ", err)
		return false
	}
	return true
}
