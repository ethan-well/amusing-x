CREATE TABLE `books`
(
    `id`          bigint                           NOT NULL AUTO_INCREMENT,
    `book_name`   varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'book name',
    `author`      varchar(255) COLLATE utf8mb4_bin NOT NULL DEFAULT '' COMMENT 'author name',
    `create_time` datetime                         NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime                         NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;