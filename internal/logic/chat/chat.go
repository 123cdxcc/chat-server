package chat

import (
	"context"
	"fmt"
	v1 "im-chat/api/ws/v1"
	"im-chat/internal/dao"
	"im-chat/internal/model"
	"im-chat/internal/model/do"
	"im-chat/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

type ChannelManager struct {
	channels   map[int64]*model.ChatChannel
	handleChan chan *v1.Message
}

func NewChannelManager() *ChannelManager {
	manager := &ChannelManager{
		channels:   make(map[int64]*model.ChatChannel),
		handleChan: make(chan *v1.Message, 100),
	}
	manager.Start()
	return manager
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
		c.channels[userID] = model.NewChatChannel(userID)
		_, err := dao.User.Ctx(ctx).Where("id = ?", userID).Data(do.User{
			Online: true,
		}).Update()
		if err != nil {
			glog.Warningf(ctx, "更新用户[%d]上线状态失败: %v", userID, err)
		}
		glog.Debugf(ctx, "用户[%d]已上线", userID)
	}
	c.channels[userID].AddConnection(ws)
}

func (c *ChannelManager) RemoveChannel(ctx context.Context, userID int64, ws *websocket.Conn) {
	channel, ok := c.channels[userID]
	if !ok {
		return
	}
	channel.RemoveConnection(ws)
	if len(channel.ChannelConnections) == 0 {
		delete(c.channels, userID)
		glog.Debugf(ctx, "用户[%d]已离线", userID)
		_, err := dao.User.Ctx(ctx).Where("id = ?", userID).Data(do.User{
			Online: false,
		}).Update()
		if err != nil {
			glog.Warningf(ctx, "更新用户[%d]离线状态失败: %v", userID, err)
		}
	}
}

func (c *ChannelManager) SendUserMessage(ctx context.Context, userID int64, data *v1.Message) error {
	channel, err := c.GetChannel(userID)
	if err != nil {
		if !gerror.HasCode(err, gcode.CodeNotFound) {
			return err
		}
		glog.Debugf(ctx, "用户[%d]无在线设备", userID)
		return nil
	}
	return channel.SendMessage(data)
}

func (c *ChannelManager) SendUsersMessage(ctx context.Context, userIDs []int64, data *v1.Message) error {
	for _, userID := range userIDs {
		err := c.SendUserMessage(ctx, userID, data)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *ChannelManager) Start() {
	go func() {
		glog.Debugf(context.Background(), "message handle task started")
		for message := range c.handleChan {
			c.handleMessage(context.Background(), message)
		}
	}()
}

func (c *ChannelManager) Stop() {
	close(c.handleChan)
}

func (c *ChannelManager) HandleMessage(message *v1.Message) {
	c.handleChan <- message
}

// 处理消息
func (c *ChannelManager) handleMessage(ctx context.Context, message *v1.Message) {
	switch message.MessageType {
	case v1.MessageTypeHeartbeat:
		return
	case v1.MessageTypeChatData:
		data := message.Data.(v1.ChatData)
		if data.To == nil {
			glog.Debugf(ctx, "消息[%s]没有目标房间", data.ID)
			return
		}
		data.ID = utility.NewID()
		roomCol := dao.UserRoomRelation.Columns()
		vals, err := dao.UserRoomRelation.Ctx(ctx).
			InnerJoin(dao.User.Table(), fmt.Sprintf("%s.%s = %s.%s", dao.User.Table(), dao.User.Columns().Id, dao.UserRoomRelation.Table(), dao.UserRoomRelation.Columns().UserId)).
			Where("room_id = ?", data.To.Id).
			Where("user.online = ?", true).
			Fields([]string{roomCol.UserId}).
			Array()
		if err != nil {
			glog.Warningf(ctx, "消息[%s]获取房间用户失败", data.ID)
			return
		}
		if len(vals) == 0 {
			glog.Debugf(ctx, "消息[%s]没有目标用户", data.ID)
			return
		}
		userIDs := gconv.Int64s(vals)
		err = c.SendUsersMessage(ctx, userIDs, &v1.Message{
			MessageType: v1.MessageTypeChatData,
			Data:        data,
		})
		if err != nil {
			glog.Warningf(ctx, "消息[%s]发送失败", data.ID)
		}
	}
}
