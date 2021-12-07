CREATE TABLE `oauth_info` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_id` bigint NOT NULL DEFAULT '0' COMMENT 'user id',
    `provider` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'provider',
    `outer_id` bigint NOT NULL DEFAULT '0' COMMENT 'outer uer id',
    `login` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'outer login name',
    `avatar_url` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'outer avatar url',
    `email` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'outer email',
    `access_token` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'access_token',
    `code` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'code',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;