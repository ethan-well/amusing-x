CREATE TABLE `action` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `action` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'action',
                          `desc` bigint NOT NULL DEFAULT '0' COMMENT 'action desc',
                          `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                          `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          UNIQUE KEY `uk_action` (`action`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;