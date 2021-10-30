create table books (
    id bigint not null auto increment,
    book_name varchar(255) not null default '' commit "book name",
    author varchar(255) not null default '' commit "author name",
    create_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    primary key id
)