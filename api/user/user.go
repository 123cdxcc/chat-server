// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"im-chat/api/user/v1"
)

type IUserV1 interface {
	GetUserDetail(ctx context.Context, req *v1.GetUserDetailReq) (res *v1.GetUserDetailRes, err error)
}
