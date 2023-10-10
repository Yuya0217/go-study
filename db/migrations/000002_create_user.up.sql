CREATE TABLE IF NOT EXISTS `users` ( 
    `id` bigint(20) NOT NULL AUTO_INCREMENT
    , `password` varchar (255) NOT NULL
    , `name` varchar (255) NOT NULL
    , `email` varchar (255) NOT NULL
    , `created_at` datetime NOT NULL DEFAULT current_timestamp
    , `updated_at` datetime NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
    , `deleted_at` datetime DEFAULT NULL
    , PRIMARY KEY (`id`)
    , UNIQUE KEY `index_users_on_email` (`email`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;