-- Create categories table
CREATE TABLE `categories` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` nvarchar(128) NOT NULL
);

-- Create orders table
CREATE TABLE `orders` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL,
  `number_of_product` int NOT NULL
);

-- Create products table
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

-- Create order_details table
CREATE TABLE `order_details` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL,
  `quantity` int NOT NULL,
  FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

-- Create product_images table
CREATE TABLE `product_images` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `image_url` varchar(255) NOT NULL,
  `product_id` int NOT NULL,
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`)
);

-- Create option_groups table
CREATE TABLE `option_groups` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` nvarchar(255) NOT NULL,
  `max_selected_options` int DEFAULT 1 NOT NULL,
  `min_selected_options` int DEFAULT 1 NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL
);

-- Create options table
CREATE TABLE `options` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `name` nvarchar(128) NOT NULL,
  `description` longtext,
  `group_id` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL,
  FOREIGN KEY (`group_id`) REFERENCES `option_groups` (`id`)
);

-- Create options_products table
CREATE TABLE `options_products` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `price_modifier_value` float,
  `product_id` int NOT NULL,
  `option_id` int NOT NULL,
  `created_at` datetime(6) NOT NULL,
  `last_updated_at` datetime(6) NOT NULL,
  FOREIGN KEY (`product_id`) REFERENCES `products` (`id`),
  FOREIGN KEY (`option_id`) REFERENCES `options` (`id`)
);

-- Create selected_options_details table
CREATE TABLE `selected_options_details` (
  `id` int PRIMARY KEY AUTO_INCREMENT,
  `detail_id` int NOT NULL,
  `option_product_id` int NOT NULL,
  FOREIGN KEY (`detail_id`) REFERENCES `order_details` (`id`),
  FOREIGN KEY (`option_product_id`) REFERENCES `options_products` (`id`)
);

-- Adding indexes for performance optimization
CREATE INDEX `idx_order_details_order_id` ON `order_details` (`order_id`);
CREATE INDEX `idx_order_details_product_id` ON `order_details` (`product_id`);
CREATE INDEX `idx_product_images_product_id` ON `product_images` (`product_id`);
CREATE INDEX `idx_products_category_id` ON `products` (`category_id`);
CREATE INDEX `idx_options_group_id` ON `options` (`group_id`);
CREATE INDEX `idx_options_products_product_id` ON `options_products` (`product_id`);
CREATE INDEX `idx_options_products_option_id` ON `options_products` (`option_id`);
CREATE INDEX `idx_selected_options_details_detail_id` ON `selected_options_details` (`detail_id`);
CREATE INDEX `idx_selected_options_details_option_product_id` ON `selected_options_details` (`option_product_id`);
