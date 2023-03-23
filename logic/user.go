package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/jwt"
	"blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户名存不存在
	if err := mysql.CheckUerExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userId := snowflake.GenID()
	u := &models.User{
		UserID:   userId,
		Username: p.Username,
		Password: p.Password,
	}
	// 3.密码加密
	// 4.保存到数据库
	return mysql.InsertUser(u)
}

func Login(p *models.ParamLogin) (token string, err error) {

	user := &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	if err := mysql.Login(user); err != nil {
		return "", err
	}

	// 生成JWT Token
	return jwt.GenToken(user.UserID)

}
