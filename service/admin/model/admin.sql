CREATE TABLE `admin_user` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `user_name` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户名称',
    `password` varchar(255)  NOT NULL DEFAULT '' COMMENT '用户密码',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 ;