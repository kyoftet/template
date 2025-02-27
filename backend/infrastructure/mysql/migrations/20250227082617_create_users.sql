-- Create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `user_id` varchar(255) NOT NULL, `name` varchar(255) NOT NULL, `password` varchar(255) NOT NULL, `mail_address` varchar(255) NOT NULL, `introduction` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `user_id` (`user_id`), UNIQUE INDEX `mail_address` (`mail_address`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
