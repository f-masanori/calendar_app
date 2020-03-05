
-- +migrate Up
CREATE TABLE IF NOT EXISTS events (
  id int(11) unsigned not null auto_increment,
  uid varchar(255) not null,
  date INT UNSIGNED not NULL,
  event varchar(255),
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);
-- +migrate Down
DROP TABLE IF EXISTS events;