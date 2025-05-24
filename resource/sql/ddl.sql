-- 版本: 1.0.0

-- 用户表
CREATE TABLE `user` (
                        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                        `username` VARCHAR(255) NOT NULL UNIQUE COMMENT '用户名',
                        PRIMARY KEY (`id`)
) COMMENT='用户表';

-- 房间表
CREATE TABLE `room` (
                        `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '房间ID',
                        `name` VARCHAR(255) NOT NULL COMMENT '房间名称',
                        PRIMARY KEY (`id`)
) COMMENT='房间表';

-- 用户与房间关系表
CREATE TABLE `user_room_relation` (
                                      `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '关系ID',
                                      `user_id` BIGINT NOT NULL COMMENT '用户ID',
                                      `room_id` BIGINT NOT NULL COMMENT '房间ID',
                                      `role` VARCHAR(50) DEFAULT NULL COMMENT '用户在房间中的角色',
                                      `subscribed` BOOLEAN DEFAULT TRUE COMMENT '是否订阅消息',
                                      `joined_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '加入时间',
                                      PRIMARY KEY (`id`),
                                      UNIQUE KEY `uniq_user_room` (`user_id`, `room_id`),
                                      KEY `idx_user_id` (`user_id`),
                                      KEY `idx_room_id` (`room_id`),
                                      KEY `idx_room_joined_at` (`room_id`, `joined_at`)
) COMMENT='用户与房间关系表';

-- 版本: 1.0.1

-- 用户增加是否在线标识
ALTER TABLE `user` ADD COLUMN `online` BOOLEAN DEFAULT FALSE COMMENT '是否在线';

-- 版本: 1.0.2

-- 用户与用户好友关系表
CREATE TABLE `user_friend_relation` (
    `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '关系ID',
    `user_id` BIGINT NOT NULL COMMENT '用户ID',
    `friend_id` BIGINT NOT NULL COMMENT '好友ID',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_user_friend` (`user_id`, `friend_id`)
) COMMENT='用户与用户好友关系表';

-- 聊天用户记录表，废弃原因：采用视图
-- CREATE TABLE `chat_record` (
--     `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT '记录ID',
--     `type` VARCHAR(50) NOT NULL COMMENT '聊天类型(user/room)',
--     `object_id` BIGINT NOT NULL COMMENT '聊天对象ID',
--     `last_message_id` BIGINT NOT NULL COMMENT '最后一条消息ID',
--     `last_message_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '最后消息时间',
--     PRIMARY KEY (`id`),
--     KEY `idx_type_object_id` (`type`, `object_id`),
--     KEY `idx_last_message_time` (`last_message_time`)
-- )

-- 聊天记录表
CREATE TABLE `chat_message` (
    `id` VARCHAR(255) NOT NULL COMMENT '消息ID',
    `client_seq_id` VARCHAR(255) NOT NULL COMMENT '客户端序列号',
    `sender_id` BIGINT NOT NULL COMMENT '发送者ID(用户ID)',
    `receiver_id` BIGINT NOT NULL COMMENT '接收者ID(用户ID或房间ID, 根据receiver_type确定)',
    `receiver_type` VARCHAR(50) NOT NULL COMMENT '接收者类型(user/room)',
    `content` TEXT NOT NULL COMMENT '消息内容',
    `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) COMMENT='聊天记录表';