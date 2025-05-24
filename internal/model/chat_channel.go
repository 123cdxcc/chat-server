package model

import (
	"context"
	v1 "im-chat/api/chat/v1"
	"time"

	"slices"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"
)

type ChatChannel struct {
	UserID             int64
	ChannelConnections []*websocket.Conn
	sendChannel        chan *v1.Message
	keepAliveTicker    *time.Ticker
	closeChan          chan struct{}
}

func NewChatChannel(userID int64) *ChatChannel {
	c := &ChatChannel{
		UserID:             userID,
		ChannelConnections: make([]*websocket.Conn, 0, 1),
		sendChannel:        make(chan *v1.Message, 100),
		keepAliveTicker:    time.NewTicker(10 * time.Second),
		closeChan:          make(chan struct{}),
	}
	c.Start()
	return c
}

func (c *ChatChannel) AddConnection(conn *websocket.Conn) {
	c.ChannelConnections = append(c.ChannelConnections, conn)
}

func (c *ChatChannel) websocketEqual(ws1, ws2 *websocket.Conn) bool {
	return ws1.LocalAddr().String() == ws2.LocalAddr().String() && ws1.RemoteAddr().String() == ws2.RemoteAddr().String()
}

func (c *ChatChannel) RemoveConnection(conn *websocket.Conn) {
	for i, connection := range c.ChannelConnections {
		if c.websocketEqual(connection, conn) {
			c.ChannelConnections = slices.Delete(c.ChannelConnections, i, i+1)
			break
		}
	}
}

func (c *ChatChannel) SendMessage(message *v1.Message) error {
	c.sendChannel <- message
	return nil
}

// 后台发送消息任务
func (c *ChatChannel) runBackgroundSendMessageTask() {
	go func() {
		for message := range c.sendChannel {
			for _, conn := range c.ChannelConnections {
				err := conn.WriteJSON(message)
				if err != nil {
					glog.Warningf(context.Background(), "用户[%d]设备[%s->%s]发送消息[%v]失败", c.UserID, conn.LocalAddr().String(), conn.RemoteAddr().String(), message.Data.(v1.ChatDataOutput).ID)
				}
			}
		}
	}()
}

// 自动保活
func (c *ChatChannel) runAutoKeepAliveTask() {
	go func() {
		for {
			select {
			case <-c.keepAliveTicker.C:
				c.SendMessage(&v1.Message{
					Type: v1.MessageTypeHeartbeat,
					Data: "keep alive",
				})
			case <-c.closeChan:
				c.keepAliveTicker.Stop()
				return
			}
		}
	}()
}

func (c *ChatChannel) Start() {
	c.runBackgroundSendMessageTask()
	c.runAutoKeepAliveTask()
}

func (c *ChatChannel) Stop() {
	close(c.closeChan)
	close(c.sendChannel)
}
