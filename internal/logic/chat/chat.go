package chat

import (
	"context"
	"fmt"
	v1 "im-chat/api/chat/v1"
	"im-chat/internal/dao"
	"im-chat/internal/model"
	"im-chat/internal/model/do"
	"im-chat/internal/model/entity"
	"im-chat/utility"
	"im-chat/utility/auth"
	"sync"

	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

type ChannelManager struct {
	channels   *sync.Map
	handleChan chan *v1.Message
}

func NewChannelManager() *ChannelManager {
	manager := &ChannelManager{
		channels:   &sync.Map{},
		handleChan: make(chan *v1.Message, 100),
	}
	manager.start()
	return manager
}

func (c *ChannelManager) GetChannel(userID int64) (*model.ChatChannel, bool) {
	channel, ok := c.channels.Load(userID)
	if !ok {
		return nil, false
	}
	return channel.(*model.ChatChannel), true
}

func (c *ChannelManager) AddChannel(ctx context.Context, userID int64, ws *websocket.Conn) {
	channel, ok := c.channels.Load(userID)
	if !ok {
		channel = model.NewChatChannel(userID)
		c.channels.Store(userID, channel)
		_, err := dao.User.Ctx(ctx).Where("id = ?", userID).Data(do.User{
			Online: true,
		}).Update()
		if err != nil {
			glog.Warningf(ctx, "更新用户[%d]上线状态失败: %v", userID, err)
		}
		glog.Debugf(ctx, "用户[%d]已上线", userID)
	}
	channel.(*model.ChatChannel).AddConnection(ws)
}

func (c *ChannelManager) RemoveChannel(ctx context.Context, userID int64, ws *websocket.Conn) {
	channel, ok := c.GetChannel(userID)
	if !ok {
		return
	}
	channel.RemoveConnection(ws)
	if len(channel.ChannelConnections) == 0 {
		c.channels.Delete(userID)
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
	channel, ok := c.GetChannel(userID)
	if !ok {
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

func (c *ChannelManager) start() {
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

func (c *ChannelManager) HandleMessage(ctx context.Context, message *v1.Message) {
	if message.Type == v1.MessageTypeChatData {
		userID := auth.GetSessionUserID(ctx)
		if userID == 0 {
			glog.Warningf(ctx, "用户未登录")
			return
		}
		data := safeGetMessageChatData(message)
		if data == nil {
			glog.Warningf(ctx, "消息转换失败: %v", message)
			return
		}
		data.SenderID = userID
		message.Data = data
	}
	c.handleChan <- message
}

// 获取房间在线用户
func (c *ChannelManager) getRoomOnlineUsers(ctx context.Context, roomID int64) ([]int64, error) {
	vals, err := dao.UserRoomRelation.Ctx(ctx).
		InnerJoin(dao.User.Table(), fmt.Sprintf("%s.%s = %s.%s", dao.User.Table(), dao.User.Columns().Id, dao.UserRoomRelation.Table(), dao.UserRoomRelation.Columns().UserId)).
		Where("room_id = ?", roomID).
		Where("user.online = ?", true).
		Fields([]string{dao.UserRoomRelation.Columns().UserId}).
		Array()
	if err != nil {
		return nil, err
	}
	userIDs := make([]int64, 0, len(vals))
	for _, val := range vals {
		userIDs = append(userIDs, gconv.Int64(val))
	}
	return userIDs, nil
}

func safeGetMessageChatData(message *v1.Message) *v1.ChatDataInput {
	if message.Type != v1.MessageTypeChatData {
		return nil
	}
	var data *v1.ChatDataInput
	err := gconv.Struct(message.Data, &data)
	if err != nil {
		return nil
	}
	return data
}

// 处理消息
func (c *ChannelManager) handleMessage(ctx context.Context, message *v1.Message) {
	switch message.Type {
	case v1.MessageTypeHeartbeat:
		return
	case v1.MessageTypeChatData:
		data := safeGetMessageChatData(message)
		if data == nil {
			glog.Warningf(ctx, "消息转换失败: %v", message)
			return
		}
		if data.Receiver == nil {
			glog.Debugf(ctx, "消息[%s]没有接受者", data.ClientSeqID)
			return
		}
		switch data.Receiver.Type {
		case v1.ChatObjectTypeUser:
			// 发送消息用户信息
			sender := new(entity.User)
			err := dao.User.Ctx(ctx).WherePri(data.SenderID).Scan(sender)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]获取发送者信息失败: %v", data.ClientSeqID, err)
				return
			}
			// 接收消息用户信息
			receiver := new(entity.User)
			err = dao.User.Ctx(ctx).WherePri(data.Receiver.ID).Scan(receiver)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]获取用户信息失败: %v", data.ClientSeqID, err)
				return
			}
			data.Receiver.Name = receiver.Username
			message := &v1.ChatDataOutput{
				ID:             utility.NewUUID(),
				AckClientSeqID: data.ClientSeqID,
				Sender:         &v1.Sender{ID: sender.Id, Name: sender.Username},
				Receiver:       data.Receiver,
				Content:        data.Content,
			}
			err = c.saveMessage(ctx, message)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]保存失败: %v", message.ID, err)
				return
			}
			err = c.SendUserMessage(ctx, data.Receiver.ID, &v1.Message{
				Type: v1.MessageTypeChatData,
				Data: message,
			})
			if err != nil {
				glog.Warningf(ctx, "消息[%s]发送失败: %v", message.ID, err)
			}
		case v1.ChatObjectTypeRoom:
			// 接收消息的房间信息
			room := new(entity.Room)
			err := dao.Room.Ctx(ctx).WherePri(data.Receiver.ID).Scan(room)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]获取房间信息失败: %v", data.ClientSeqID, err)
				return
			}
			data.Receiver.Name = room.Name
			// 获取房间在线用户
			userIDs, err := c.getRoomOnlineUsers(ctx, data.Receiver.ID)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]获取房间用户失败: %v", data.ClientSeqID, err)
				return
			}
			if len(userIDs) == 0 {
				glog.Debugf(ctx, "消息[%s]没有目标用户", data.ClientSeqID)
				return
			}
			sender := new(entity.User)
			err = dao.User.Ctx(ctx).WherePri(data.SenderID).Scan(sender)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]获取发送者信息失败: %v", data.ClientSeqID, err)
				return
			}
			message := &v1.ChatDataOutput{
				ID:             utility.NewUUID(),
				AckClientSeqID: data.ClientSeqID,
				Sender:         &v1.Sender{ID: sender.Id, Name: sender.Username},
				Receiver:       data.Receiver,
				Content:        data.Content,
			}
			err = c.saveMessage(ctx, message)
			if err != nil {
				glog.Warningf(ctx, "消息[%s]保存失败: %v", message.ID, err)
				return
			}
			err = c.SendUsersMessage(ctx, userIDs, &v1.Message{
				Type: v1.MessageTypeChatData,
				Data: message,
			})
			if err != nil {
				glog.Warningf(ctx, "消息[%s]发送失败: %v", message.ID, err)
			}
		}
	}
}

func (c *ChannelManager) saveMessage(ctx context.Context, message *v1.ChatDataOutput) error {
	_, err := dao.ChatMessage.Ctx(ctx).Data(do.ChatMessage{
		Id:           message.ID,
		ClientSeqId:  message.AckClientSeqID,
		SenderId:     message.Sender.ID,
		ReceiverId:   message.Receiver.ID,
		ReceiverType: message.Receiver.Type,
		Content:      message.Content,
	}).Insert()
	if err != nil {
		glog.Warningf(ctx, "消息[%s]保存失败: %v", message.ID, err)
		return err
	}
	return nil
}
