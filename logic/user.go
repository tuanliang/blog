package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
	"fmt"
)

func SignUp(p *models.ParamSignUp) (err error) {
	fmt.Println(11111)
	// 1.判断用户存不存在
	if err = mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	// 2.生成uid
	userID := snowflake.GenID()
	// 构造user实例
	u := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	// 密码加密
	// 3.保存进数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	// 传递的是指针，可以拿到userID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}

	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}
