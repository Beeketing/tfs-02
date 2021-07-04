CREATE TABLE `order` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` text CHARACTER SET latin1,
  `currency` text CHARACTER SET latin1 NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `processed_at` datetime DEFAULT NULL,
  `cancelled_at` datetime DEFAULT NULL,
  `cart_token` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `total_price` double DEFAULT NULL,
  `checkout_token` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `total_weight_in_gram` double DEFAULT NULL,
  `transaction_id` bigint(20) DEFAULT NULL,
  `customer_name` varchar(100) DEFAULT NULL,
  `email` text CHARACTER SET latin1,
  `client_ip` text CHARACTER SET latin1,
  `client_country_code` text CHARACTER SET latin1,
  `phone` text CHARACTER SET latin1,
  `payment_gateway` text CHARACTER SET latin1,
  `fulfilled_at` datetime DEFAULT NULL,
  `additional_information` text COLLATE utf8_unicode_ci,
  `thankyou_email_sent` tinyint(1),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;




INSERT INTO `order` (`name`, `currency`, `created_at`, `updated_at`, `processed_at`, `cancelled_at`, `cart_token`, `total_price`, `checkout_token`, `total_weight_in_gram`, `transaction_id`, `customer_name`, `email`, `client_ip`, `client_country_code`, `phone`, `payment_gateway`, `fulfilled_at`, `additional_information`) VALUES
('#sb-01', 'USD', NOW(), NOW(), NOW(), NULL, '98ddd1jsg761g8yt1ghdk', 121, '8uhfsafklluhsyavnx98hsoc0', 0.5, 738892783, 'Truong Bui Van', 'truongbui@shopbase.com', '58.167.16.18', 'VN', '+84356172989', 'Paypal express', NULL, 0);