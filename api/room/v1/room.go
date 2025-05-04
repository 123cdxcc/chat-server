package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"im-chat/internal/model/entity"
)

type GetRoomListReq struct {
	g.Meta `path:"/room" tags:"房间" method:"get" summary:"房间列表"`
}

type GetRoomListRes struct {
	List []*entity.Room `json:"list" dc:"房间列表"`
}

// 创建房间
type CreateRoomReq struct {
	g.Meta `path:"/room" tags:"房间" method:"post" summary:"创建房间"`
	Name   string `json:"name" v:"required#房间名称不能为空" dc:"房间名称"`
}

type CreateRoomRes struct {
	Room *entity.Room `json:"room" dc:"房间信息"`
}

// 加入房间
type JoinRoomReq struct {
	g.Meta `path:"/room/join" tags:"房间" method:"post" summary:"加入房间"`
	RoomId int64 `json:"room_id" v:"required#房间ID不能为空" dc:"房间ID"`
}

type JoinRoomRes struct {
	Room *entity.Room `json:"room" dc:"房间信息"`
}
