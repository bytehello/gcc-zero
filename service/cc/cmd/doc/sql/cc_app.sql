CREATE TABLE `cc_app` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `app_key` varchar(64) NOT NULL DEFAULT '' COMMENT 'app key',
    `app_name` varchar(255)  NOT NULL DEFAULT '' COMMENT 'app名称',
    `desc` varchar(244) NOT NULL DEFAULT '' COMMENT 'app描述',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time` timestamp null default null,
    PRIMARY KEY (`id`),UNIQUE INDEX uidx_app_key(app_key)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT 'app表';