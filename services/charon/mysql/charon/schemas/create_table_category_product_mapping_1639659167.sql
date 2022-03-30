CREATE TABLE `category_product_mapping` (
                                            `id` bigint NOT NULL AUTO_INCREMENT,
                                            `category_id` bigint NOT NULL DEFAULT '0' COMMENT 'user id',
                                            `product_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'category name',
                                            `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;