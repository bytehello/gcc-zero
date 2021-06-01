CREATE TABLE `cc_kv`
(
    `id`              bigint       NOT NULL AUTO_INCREMENT,
    `key`             varchar(255) NOT NULL DEFAULT '' COMMENT '键名',
    `value`           text COMMENT '值',
    `version`         int          not null default 0 comment '版本',
    `create_revision` int          not null default 0 comment 'create_revision',
    `mod_revision`    int          not null default 0 comment 'mod_revision',
    `create_time`     timestamp    NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time`     timestamp    NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time`    timestamp    null     default null,
    PRIMARY KEY (`id`),
    index idx_key (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT 'kv表';