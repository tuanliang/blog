package controller

import "blog/models"

// 专门用来放接口文档用到的model

type _ResponsePostList struct {
	Code    ResCode                 `json:"code"`
	Message string                  `json:"message"`
	Data    []*models.ApiPostDetail `json:"data"`
}
