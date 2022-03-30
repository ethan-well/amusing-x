CREATE TABLE `product_stock` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `sub_product_id` bigint NOT NULL DEFAULT '0' COMMENT 'sub product id',
                                 `value` bigint NOT NULL DEFAULT '0' COMMENT 'sub product stock',
                                 `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`),
                                 UNIQUE KEY `uk_sub_product_id` (`sub_product_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;