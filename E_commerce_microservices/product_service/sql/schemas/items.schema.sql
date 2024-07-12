CREATE TABLE items (
  id bigint auto_increment primary key,
  name varchar(128),
  category_id int ,
  short_description varchar(255),
  original_price float,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp ON UPDATE current_timestamp
);  
