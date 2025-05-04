// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package room

import (
	"context"

	"im-chat/api/room/v1"
)

type IRoomV1 interface {
	GetRoomList(ctx context.Context, req *v1.GetRoomListReq) (res *v1.GetRoomListRes, err error)
	CreateRoom(ctx context.Context, req *v1.CreateRoomReq) (res *v1.CreateRoomRes, err error)
	JoinRoom(ctx context.Context, req *v1.JoinRoomReq) (res *v1.JoinRoomRes, err error)
}
