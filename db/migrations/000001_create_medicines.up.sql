CREATE TABLE IF NOT EXISTS `medicines` ( 
    `id` bigint(20) NOT NULL AUTO_INCREMENT
    , `medicine_name` varchar (255) NOT NULL
    , `medicine_type` varchar (255) NOT NULL
    , `created_at` datetime NOT NULL DEFAULT current_timestamp
    , `updated_at` datetime NOT NULL DEFAULT current_timestamp ON UPDATE current_timestamp
    , `deleted_at` datetime DEFAULT NULL
    , PRIMARY KEY (`id`)
) DEFAULT CHARSET = utf8 COLLATE = utf8_bin;