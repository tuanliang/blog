package logic

import (
	"blog/dao/mysql"
	"blog/models"
)

func GetCommunityList() ([]*models.Community, error) {
	// 查找所有的community，并返回
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
