package logic

import (
	"blog/dao/mysql"
	"blog/models"
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
