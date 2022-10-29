-- +goose Up
CREATE
DATABASE IF NOT EXISTS v_pay_orders;

-- CREATE TABLE IF NOT EXISTS `orders` (
--                           `order_no` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
--                           `merchant_id` longtext,
--                           `app_id` int(11) DEFAULT NULL,
--                           `status` int(11) DEFAULT NULL,
--                           `amount` bigint(20) DEFAULT NULL,
--                           `product_code` longtext,
--                           `description` longtext,
--                           `create_time` datetime(3) DEFAULT NULL,
--                           `title` longtext,
--                           PRIMARY KEY (`order_no`)
-- ) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;
