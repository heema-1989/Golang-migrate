package controllers

import (
	"chat-application/initializers"
	"chat-application/models"
	users "chat-application/sqlc-models"
	"chat-application/utils"
	"context"
	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/satori/go.uuid"
	"time"
)

type RegisterController struct {
	beego.Controller
}

var (
	user      models.Token
	secretKey []byte
)

// URLMapping ...
func URLMapping(mapping *RegisterController) {
	mapping.Mapping("RegisterUser", mapping.RegisterUser)
	mapping.Mapping("VerifyUser", mapping.VerifyUser)
	mapping.Mapping("VerifiedUser", mapping.VerifiedUser)
}

// RegisterUser ...
// @Title Get
// @Description get Register Page
// @Success 200 {object} models.Token
// @Failure 403
// @router / [get]
func (register *RegisterController) RegisterUser() {
	register.TplName = "default/register.html"
}

// VerifyUser ...
// @Title  Post
// @Description get verify-user page
// @Success 200 {object} models.Token
// @Failure 403
// @router / [post]
func (register *RegisterController) VerifyUser() {
	user = register.FormParser(user)
	logs.Info(user.Email)
	if !UserAlreadyExist(user) {
		user.VerifyID = uuid.NewV4()
		user.JwtToken, secretKey = utils.GenerateJWTToken(time.Now().Add(1*time.Minute), &user)
		register.SetUserTemplateNames(user)
		utils.SendMail(user.Email, user)
		register.TplName = "default/verify-register.html"
		logs.Info(user)

	} else {
		register.SetUserTemplateNames(user)
		register.TplName = "default/user-exists.html"
	}
}

// VerifiedUser ...
// @Title  Get
// @Description get verify-user page
// @Success 200 {object} models.Token
// @Failure 403
// @router /verified [post]
func (register *RegisterController) VerifiedUser() {
	var (
		flag bool = true
	)
	register.SetUserTemplateNames(user)
	err := utils.ValidateJwt(user, user.JwtToken, secretKey)
	if err != nil {
		flag = false
		register.Data["Error"] = flag
		register.TplName = "default/welcome.html"
		return
	}
	register.Data["Error"] = flag
	register.TplName = "default/welcome.html"
	result, createErr := initializers.Db.CreateUser(context.Background(), users.CreateUserParams{
		FullName: user.FullName,
		UserName: user.UserName,
		Email:    user.Email,
		Password: user.Password,
		VerifyID: user.VerifyID,
	})
	utils.CheckError(createErr, "Error creating record")
	logs.Info("User successfully created: ", result)
}

func (register *RegisterController) SetUserTemplateNames(user models.Token) {
	register.Data["VerifyId"] = user.VerifyID
	register.Data["UserName"] = user.UserName
	register.Data["FullName"] = user.FullName
	register.Data["Email"] = user.Email
	register.Data["Password"] = user.Password
}

func (register *RegisterController) FormParser(user models.Token) models.Token {
	parseErr := register.ParseForm(&user)
	utils.CheckError(parseErr, "Error parsing form")
	if register.Ctx.Input.Method() == "POST" {
		user.FullName = register.GetString("fullname")
		user.UserName = register.GetString("username")
		user.Email = register.GetString("email")
		user.Password = register.GetString("password")
		return user
	}
	return models.Token{}
}

func UserAlreadyExist(user models.Token) bool {
	_, err := initializers.Db.GetUserEmail(context.Background(), user.Email)
	if err != nil {
		return false
	}
	return true
}
