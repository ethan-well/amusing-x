CREATE TABLE `product_image` (
                                 `id` bigint NOT NULL AUTO_INCREMENT,
                                 `product_id` bigint NOT NULL DEFAULT '0' COMMENT 'product id',
                                 `product_level` bigint NOT NULL DEFAULT '0' COMMENT 'product level',
                                 `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'image url',
                                 `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0' COMMENT 'image title',
                                 `uploader_type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT 'local' COMMENT 'uploader type',
                                 `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                 `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;