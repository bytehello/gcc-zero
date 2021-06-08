CREATE TABLE `cc_cluster`
(
    `id`           bigint       NOT NULL AUTO_INCREMENT,
    `app_id`       bigint       NOT NULL default '' comment 'cc_app id',
    `cluster_name` varchar(64)  NOT NULL DEFAULT '' COMMENT '集群名',
    `desc`         varchar(244) NOT NULL DEFAULT '' COMMENT 'a描述',
    `create_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time`  timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time` timestamp    null     default null,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '集群表';