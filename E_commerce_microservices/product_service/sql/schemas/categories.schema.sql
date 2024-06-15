create table categories (
  id int not null auto_increment primary key,
  name varchar(128) not null unique,
  created_at datetime default current_timestamp,
  updated_at datetime default current_timestamp
)
