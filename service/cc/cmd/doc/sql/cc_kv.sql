CREATE TABLE `cc_kv`
(
    `id`              bigint           NOT NULL AUTO_INCREMENT,
    `key`             varchar(255)     NOT NULL DEFAULT '' COMMENT '键名',
    `value`           text COMMENT '值',
    `version`         int              not null default 0 comment '版本',
    `create_revision` int              not null default 0 comment 'create_revision',
    `mod_revision`    int              not null default 0 comment 'mod_revision',
    `push_status`     tinyint unsigned not null default 0 comment '推送状态0未推送1已经推送',
    `create_time`     timestamp        NULL     DEFAULT CURRENT_TIMESTAMP,
    `update_time`     timestamp        NULL     DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_time`    timestamp        null     default null,
    `pushed_time`     timestamp        null     default null comment '推送时间',
    PRIMARY KEY (`id`),
    index idx_key (`key`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT 'kv表';