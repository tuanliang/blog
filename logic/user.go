package logic

import (
	"blog/dao/mysql"
	"blog/models"
	"blog/pkg/snowflake"
)

func SignUp(p *models.ParamSignUp) {
	// 1.判断用户名存不存在
	mysql.QueryUserByUserName()
	// 2.生成UID
	snowflake.GenID()
	// 3.密码加密
	// 4.保存到数据库
	mysql.InsertUser()
}
