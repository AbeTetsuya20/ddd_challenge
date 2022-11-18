USE app;

CREATE TABLE
    `user` (
               `user_id` VARCHAR(10) NOT NULL COMMENT 'ユーザー ID',
               `name` VARCHAR(10) NOT NULL COMMENT 'ユーザー名',
               `password` VARCHAR(10) NOT NULL COMMENT 'パスワード',
               `created_at` DATETIME NOT NULL COMMENT '作成日時',
               `updated_at` DATETIME NOT NULL COMMENT '更新日時',
               PRIMARY KEY (`user_id`),
               INDEX `user_updated_at` (`updated_at`)
) COMMENT = 'ユーザー';

CREATE TABLE
    `message` (
               `message_id` VARCHAR(10) NOT NULL COMMENT 'メッセージ ID',
               `messageTitle` VARCHAR(30) NOT NULL COMMENT 'メッセージタイトル',
               `messageBody` VARCHAR(200) NOT NULL COMMENT 'メッセージボディ',
               `author` VARCHAR(30) NOT NULL COMMENT 'メッセージ作者',
               `channel` VARCHAR(30) NOT NULL COMMENT 'チャンネル',
               `is_Send` boolean NOT NULL COMMENT '送信状況',
               `send_at` DATETIME NOT NULL COMMENT '送信日時',
               `created_at` DATETIME NOT NULL COMMENT '作成日時',
               `updated_at` DATETIME NOT NULL COMMENT '更新日時',
               PRIMARY KEY (`message_id`),
               FOREIGN KEY (`author`) REFERENCES `user` (`user_id`),
               FOREIGN KEY (`channel`) REFERENCES `channel` (`channel_id`),
               INDEX `message_updated_at` (`updated_at`)
) COMMENT = 'メッセージ';

CREATE TABLE
    `channel` (
                `channel_id` VARCHAR(10) NOT NULL COMMENT 'チャンネル ID',
                `name` VARCHAR(30) COMMENT 'チャンネル名',
                `created_at` DATETIME NOT NULL COMMENT '作成日時',
                `updated_at` DATETIME NOT NULL COMMENT '更新日時',
                PRIMARY KEY (`channel_id`),
                INDEX `channel_updated_at` (`updated_at`)
) COMMENT = 'チャンネル';

CREATE TABLE
    `joinChannelToUser` (
                  `user_id` VARCHAR(10) NOT NULL COMMENT 'ユーザー ID',
                  `user_name` VARCHAR(30) COMMENT 'ユーザー名',
                  `channel_id` VARCHAR(10) NOT NULL COMMENT 'チャンネル ID',
                  `name` VARCHAR(30) COMMENT 'チャンネル名',
                  `created_at` DATETIME NOT NULL COMMENT '作成日時',
                  `updated_at` DATETIME NOT NULL COMMENT '更新日時',
                  PRIMARY KEY (`user_id`, `channel_id`),
                  FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
                  FOREIGN KEY (`channel_id`) REFERENCES `channel` (`channel_id`),
                  INDEX `join_updated_at` (`updated_at`)
) COMMENT = 'チャンネルとユーザーの中間テーブル';
