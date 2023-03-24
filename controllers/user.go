package controllers

import (
	"blog/dao/mysql"
	"blog/logic"
	"blog/models"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	// 1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(&p); err != nil { // shouldbindjson用来做参数类型校验string，int
		zap.L().Error("SignUpHandler with invaild params：", zap.Error(err))
		// 判单err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	// 对参数进行非空等校验
	// if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.Password != p.RePassword {
	// 	zap.L().Error("SignUpHandler with invaild params")
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"msg": "请求参数有误",
	// 	})
	// 	return
	// }

	// 2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("logic.SignUp failed:", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	// 3.返回相应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {

	// 1.获取请求参数和参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(&p); err != nil { // shouldbindjson用来做参数类型校验string，int
		zap.L().Error("Login with invaild params：", zap.Error(err))
		// 判单err是不是validator.ValidationErrors类型
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))

		return
	}
	// 2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login failed：", zap.String("username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNotExist)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	// 3.返回相应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID),
		"user_name": user.Username,
		"token":     user.Token,
	})
}
