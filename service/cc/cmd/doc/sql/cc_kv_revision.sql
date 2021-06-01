CREATE TABLE `cc_kv_revision`
(
    `id`              bigint       NOT NULL AUTO_INCREMENT,
    `kv_id`           bigint       not null default 0 comment 'cc_kv主键',
    `key`             varchar(255) NOT NULL DEFAULT '' COMMENT '键名',
    `value`           text COMMENT '值',
    `version`         int          not null default 0 comment '版本',
    `create_revision` int          not null default 0 comment 'create_revision',
    `mod_revision`    int          not null default 0 comment 'mod_revision',
    `create_time`     timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time`     timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time`    timestamp    null     default null,
    PRIMARY KEY (`id`),
    index idx_kv_id (`kv_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT 'kv历史表';