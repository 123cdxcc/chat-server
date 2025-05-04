package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"net/http"
)

const userIDKey = "user-id"

func SetSessionUserID(ctx context.Context, userID int64) {
	g.RequestFromCtx(ctx).Session.Set(userIDKey, userID)
}

func GetSessionUserID(ctx context.Context) int64 {
	userID, _ := g.RequestFromCtx(ctx).Session.Get(userIDKey, 0)
	return userID.Int64()
}

func SessionAuth(r *ghttp.Request) {
	userIDVar, _ := r.Session.Get("user-id")
	if userIDVar == nil || userIDVar.String() == "" {
		r.Response.WriteStatus(http.StatusUnauthorized)
		return
	}
	r.Middleware.Next()
}
