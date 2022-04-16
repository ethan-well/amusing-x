CREATE TABLE `product_image_base64` (
                                        `id` bigint NOT NULL AUTO_INCREMENT,
                                        `base64` text COLLATE utf8mb4_bin,
                                        PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;