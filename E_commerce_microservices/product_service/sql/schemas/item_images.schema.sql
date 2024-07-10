CREATE TABLE item_images (
  id bigint auto_increment primary key,
  display_name varchar(128),
  image_url varchar(255) Not null,
  item_id bigint Not null,
  FOREIGN KEY(item_id) REFERENCES items(id)
);
