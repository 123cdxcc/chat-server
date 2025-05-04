package chat

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gorilla/websocket"
	v1 "im-chat/api/ws/v1"
	"im-chat/internal/model"
)

type ChannelManager struct {
	channels map[int64]*model.ChatChannel
}

func NewChannelManager() *ChannelManager {
	return &ChannelManager{
		channels: make(map[int64]*model.ChatChannel),
	}
}

func (c *ChannelManager) GetChannel(userID int64) (*model.ChatChannel, error) {
	channel, ok := c.channels[userID]
	if !ok {
		return nil, gerror.NewCode(gcode.CodeNotFound, "通道不存在")
	}
	return channel, nil
}

func (c *ChannelManager) AddChannel(ctx context.Context, userID int64, ws *websocket.Conn) {
	if _, ok := c.channels[userID]; !ok {
		c.channels[userID] = &model.ChatChannel{
			ChannelConnections: make([]*websocket.Conn, 0),
		}
		glog.Debugf(ctx, "用户[%d]已上线", userID)
	}
	c.channels[userID].ChannelConnections = append(c.channels[userID].ChannelConnections, ws)
}

func (c *ChannelManager) websocketEqual(ws1, ws2 *websocket.Conn) bool {
	return ws1.LocalAddr().String() == ws2.LocalAddr().String() && ws1.RemoteAddr().String() == ws2.RemoteAddr().String()
}

func (c *ChannelManager) RemoveChannel(ctx context.Context, userID int64, ws *websocket.Conn) {
	channel, ok := c.channels[userID]
	if !ok {
		return
	}
	for i, connection := range channel.ChannelConnections {
		if c.websocketEqual(ws, connection) {
			channel.ChannelConnections = append(channel.ChannelConnections[:i], channel.ChannelConnections[i+1:]...)
			break
		}
	}
	if len(channel.ChannelConnections) == 0 {
		delete(c.channels, userID)
		glog.Debugf(ctx, "用户[%d]已离线", userID)
	}
}

func (c *ChannelManager) SendUserMessage(ctx context.Context, userID int64, data *v1.ChatData) error {
	channel, err := c.GetChannel(userID)
	if err != nil {
		if !gerror.HasCode(err, gcode.CodeNotFound) {
			return err
		}
		glog.Debugf(ctx, "用户[%d]无在线设备", userID)
		return nil
	}
	for _, connection := range channel.ChannelConnections {
		err := connection.WriteJSON(data)
		if err != nil {
			glog.Warningf(ctx, "用户[%d]设备[%s->%s]发送消息失败", userID, connection.LocalAddr().String(), connection.RemoteAddr().String())
		}
	}
	return nil
}

func (c *ChannelManager) SendUsersMessage(ctx context.Context, userIDs []int64, data *v1.ChatData) error {
	for _, userID := range userIDs {
		err := c.SendUserMessage(ctx, userID, data)
		if err != nil {
			return err
		}
	}
	return nil
}
