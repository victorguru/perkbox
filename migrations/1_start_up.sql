CREATE TABLE `coupons`
(
  `id` int
(11) NOT NULL AUTO_INCREMENT,
  `name` varchar
(100) DEFAULT NULL,
  `brand` varchar
(255) DEFAULT NULL,
  `value` float DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `expiry` datetime DEFAULT NULL,
  PRIMARY KEY
(`id`)
);