CREATE TABLE `role_action` (
                               `id` bigint NOT NULL AUTO_INCREMENT,
                               `role_id` bigint NOT NULL DEFAULT '0' COMMENT 'role id',
                               `action_id` bigint NOT NULL DEFAULT '0' COMMENT 'action id',
                               `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                               `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `uk_role_action_id` (`role_id`,`action_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;