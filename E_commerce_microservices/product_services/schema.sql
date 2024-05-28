CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` nvarchar(128) NOT NULL
);
CREATE TABLE `products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` nvarchar(128) NOT NULL,
  `description` longtext,
  `original_price` float NOT NULL,
  `category_id` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL,
  FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`)
);
CREATE TABLE `product_images` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `image_url` varchar(255) NOT NULL,
  `product_id` int NOT NULL,
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);