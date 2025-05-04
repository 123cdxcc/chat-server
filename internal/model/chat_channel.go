package model

import "github.com/gorilla/websocket"

type ChatChannel struct {
	UserID             int64
	ChannelConnections []*websocket.Conn
}
