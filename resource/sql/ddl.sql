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