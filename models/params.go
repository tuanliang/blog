package models

const (
	OrderTime  = "time"
	OrderScore = "score"
)

// 定义请求的参数结构体

// ParamSignUp 注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

// ParamLogin 登录请求canshu
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 投票数据
type ParamVoteData struct {
	// UserID 从当前请求的用户中获取
	PostID    string `json:"post_id" binding:"required"`              // 帖子ID
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"` //赞成票（1）反对票（-1）取消投票（0）
}

type ParamPostList struct {
	CommunityID int64 `json:"community_id" form:"community_id"`

	Page  int64  `json:"page" form:"page"`
	Size  int64  `json:"size" form:"size"`
	Order string `json:"order" form:"order"`
}

type ParamCommunityPostList struct {
	ParamPostList
}
