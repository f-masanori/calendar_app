
-- +migrate Up
CREATE TABLE IF NOT EXISTS next_event_ids (
  id int(11) unsigned not null auto_increment,
  uid varchar(255) not null,
  next_event_id int(11) unsigned not null,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);
-- +migrate Down
DROP TABLE IF EXISTS next_event_ids;