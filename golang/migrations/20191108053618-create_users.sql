
-- +migrate Up
CREATE TABLE IF NOT EXISTS users (
    id int(11) unsigned not null auto_increment,
    uid varchar(255) not null,
    email varchar(255),
    name varchar(255),
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    primary key (id)
);
-- +migrate Down
DROP TABLE IF EXISTS users;