CREATE TABLE `book_inventory` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `book_id` bigint(20) NOT NULL DEFAULT '0' COMMENT 'book id',
    `real_inventory` bigint(20) NOT NULL DEFAULT '0' COMMENT 'real_inventory = available_inventory + locked_inventory',
    `available_inventory` bigint(20) NOT NULL DEFAULT '0' COMMENT 'available inventory',
    `locked_inventory` bigint(20) NOT NULL DEFAULT '0' COMMENT 'temporary locked inventory',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_book_id` (`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;