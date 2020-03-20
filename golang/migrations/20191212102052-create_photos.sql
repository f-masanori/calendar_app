
-- +migrate Up
CREATE TABLE IF NOT EXISTS photos (
    id int(11) unsigned not null auto_increment,
    nikki_id int unsigned not null,
    user_id int(11) unsigned not null,
    date INT UNSIGNED NOT NULL,
    photo_id int unsigned not null,
    photo varchar(255) not null,
    created_at datetime not null default current_timestamp,
    updated_at datetime not null default current_timestamp on update current_timestamp,
    primary key (id)
);
-- +migrate Down
DROP TABLE IF EXISTS photos;