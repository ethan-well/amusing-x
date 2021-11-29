create table books (
    id bigint not null AUTO_INCREMENT,
    book_name varchar(255) not null default '' COMMENT "book name",
    author varchar(255) not null default '' COMMENT "author name",
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;