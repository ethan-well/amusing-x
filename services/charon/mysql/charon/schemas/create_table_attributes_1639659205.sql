CREATE TABLE `attribute_mapping` (
                                     `id` bigint NOT NULL AUTO_INCREMENT,
                                     `attr_id` bigint NOT NULL DEFAULT '0' COMMENT 'attr id',
                                     `attr_value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'attr name',
                                     `sub_product_id` bigint NOT NULL DEFAULT '0' COMMENT 'sub product id',
                                     `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                     `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                     PRIMARY KEY (`id`),
                                     UNIQUE KEY `uk_attr_id_sub_product_id` (`attr_id`,`sub_product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;