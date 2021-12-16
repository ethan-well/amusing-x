CREATE TABLE `category` (
                              `id` bigint NOT NULL AUTO_INCREMENT,
                              `parent_id` bigint NOT NULL DEFAULT '0' COMMENT 'user id',
                              `name` varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'category name',
                              `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                              `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;