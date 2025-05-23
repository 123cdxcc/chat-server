package v1

import (
	"im-chat/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// AddFriendReq 添加好友请求
type AddFriendReq struct {
	g.Meta `path:"/friend" method:"post" tags:"好友管理" summary:"添加好友"`
	UserId int64 `json:"user_id" v:"required#好友ID不能为空" dc:"好友ID"`
}

// AddFriendRes 添加好友响应
type AddFriendRes struct{}

// GetFriendListReq 获取好友列表请求
type GetFriendListReq struct {
	g.Meta `path:"/friends" method:"get" tags:"好友管理" summary:"获取好友列表"`
}

// GetFriendListRes 获取好友列表响应
type GetFriendListRes struct {
	List []*entity.User `json:"list" dc:"好友列表"`
}
